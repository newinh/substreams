package main

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/streamingfast/cli/sflags"
	"github.com/streamingfast/substreams/client"
	"github.com/streamingfast/substreams/manifest"
	pbsubstreamsrpc "github.com/streamingfast/substreams/pb/sf/substreams/rpc/v2"
	"github.com/streamingfast/substreams/tools"
	"github.com/streamingfast/substreams/tools/test"
	"github.com/streamingfast/substreams/tui"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func init() {
	runCmd.Flags().StringP("substreams-endpoint", "e", "", "Substreams gRPC endpoint. If empty, will be replaced by the SUBSTREAMS_ENDPOINT_{network_name} environment variable, where `network_name` is determined from the substreams manifest. Some network names have default endpoints.")
	runCmd.Flags().String("substreams-api-token-envvar", "SUBSTREAMS_API_TOKEN", "name of variable containing Substreams Authentication token")
	runCmd.Flags().String("substreams-api-key-envvar", "SUBSTREAMS_API_KEY", "Name of variable containing Substreams Api Key")
	runCmd.Flags().String("network", "", "Specify the network to use for params and initialBlocks, overriding the 'network' field in the substreams package")
	runCmd.Flags().StringP("start-block", "s", "", "Start block to stream from. If empty, will be replaced by initialBlock of the first module you are streaming. If negative, will be resolved by the server relative to the chain head")
	runCmd.Flags().StringP("cursor", "c", "", "Cursor to stream from. Leave blank for no cursor")
	runCmd.Flags().StringP("stop-block", "t", "0", "Stop block to end stream at, exclusively. If the start-block is positive, a '+' prefix can indicate 'relative to start-block'")
	runCmd.Flags().Bool("final-blocks-only", false, "Only process blocks that have pass finality, to prevent any reorg and undo signal by staying further away from the chain HEAD")
	runCmd.Flags().Bool("insecure", false, "Skip certificate validation on GRPC connection")
	runCmd.Flags().Bool("plaintext", false, "Establish GRPC connection in plaintext")
	runCmd.Flags().StringP("output", "o", "", "Output mode, one of: [ui, json, jsonl, clock] Defaults to 'ui' when in a TTY is present, and 'json' otherwise")
	runCmd.Flags().StringSlice("debug-modules-initial-snapshot", nil, "List of 'store' modules from which to print the initial data snapshot (Unavailable in Production Mode)")
	runCmd.Flags().StringSlice("debug-modules-output", nil, "List of modules from which to print outputs, deltas and logs (Unavailable in Production Mode)")
	runCmd.Flags().StringSliceP("header", "H", nil, "Additional headers to be sent in the substreams request")
	runCmd.Flags().Bool("production-mode", false, "Enable Production Mode, with high-speed parallel processing")
	runCmd.Flags().Bool("skip-package-validation", false, "Do not perform any validation when reading substreams package")
	runCmd.Flags().StringArrayP("params", "p", nil, "Set a params for parameterizable modules. Can be specified multiple times. Ex: -p module1=valA -p module2=valX&valY")
	runCmd.Flags().String("test-file", "", "runs a test file")
	runCmd.Flags().Bool("test-verbose", false, "print out all the results")
	rootCmd.AddCommand(runCmd)
}

// runCmd represents the command to run substreams remotely
var runCmd = &cobra.Command{
	Use:          "run [<manifest> [<module_name>]]",
	Short:        "Stream module to standard output. Use 'substreams gui' for more tools and a better experience.",
	Long:         guiOrRunLongUsage,
	RunE:         runRun,
	Args:         cobra.RangeArgs(1, 2),
	SilenceUsage: true,
}

func runRun(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	manifestPath, outputModule, err := ruiOrGuiManifestModulePositionalParams(args)
	if err != nil {
		return err
	}

	outputMode := sflags.MustGetString(cmd, "output")

	network := sflags.MustGetString(cmd, "network")
	paramsString := sflags.MustGetStringArray(cmd, "params")
	params, err := manifest.ParseParams(paramsString)
	if err != nil {
		return fmt.Errorf("parsing params: %w", err)
	}

	readerOptions := []manifest.Option{
		manifest.WithOverrideNetwork(network),
		manifest.WithParams(params),
		manifest.WithRegistryURL(getSubstreamsRegistryEndpoint()),
	}

	if outputModule != "" {
		readerOptions = append(readerOptions, manifest.WithOverrideOutputModule(outputModule))
	}

	if sflags.MustGetBool(cmd, "skip-package-validation") {
		readerOptions = append(readerOptions, manifest.SkipPackageValidationReader())
	}

	manifestReader, err := manifest.NewReader(manifestPath, readerOptions...)
	if err != nil {
		return fmt.Errorf("manifest reader: %w", err)
	}

	pkgBundle, err := manifestReader.Read()
	if err != nil {
		return fmt.Errorf("read manifest %q: %w", manifestPath, err)
	}

	if pkgBundle == nil {
		return fmt.Errorf("no package found")
	}

	endpoint, err := manifest.ExtractNetworkEndpoint(pkgBundle.Package.Network, sflags.MustGetString(cmd, "substreams-endpoint"), zlog)
	if err != nil {
		return fmt.Errorf("extracting endpoint: %w", err)
	}

	msgDescs, err := manifest.BuildMessageDescriptors(pkgBundle.Package)
	if err != nil {
		return fmt.Errorf("building message descriptors: %w", err)
	}

	var testRunner *test.Runner
	testFile := sflags.MustGetString(cmd, "test-file")
	if testFile != "" {
		zlog.Info("running test runner", zap.String(testFile, testFile))
		testRunner, err = test.NewRunner(testFile, msgDescs, sflags.MustGetBool(cmd, "test-verbose"), zlog)
		if err != nil {
			return fmt.Errorf("failed to setup test runner: %w", err)
		}
	}

	productionMode := sflags.MustGetBool(cmd, "production-mode")
	debugModulesOutput := sflags.MustGetStringSlice(cmd, "debug-modules-output")
	if len(debugModulesOutput) == 0 {
		debugModulesOutput = nil
	}
	if debugModulesOutput != nil && productionMode {
		return fmt.Errorf("cannot set 'debug-modules-output' in 'production-mode'")
	}

	debugModulesInitialSnapshot := sflags.MustGetStringSlice(cmd, "debug-modules-initial-snapshot")
	if len(debugModulesInitialSnapshot) == 0 {
		debugModulesInitialSnapshot = nil
	}

	startBlock, readFromModule, err := readStartBlockFlag(cmd, "start-block")
	if err != nil {
		return fmt.Errorf("stop block: %w", err)
	}

	if outputModule == "" {
		mods, ok := pkgBundle.Graph.TopologicalSort()
		if ok {
			outputModule = mods[0].Name
			fmt.Printf("Selected output module: %s\n", outputModule)
		}
	}

	if readFromModule {
		sb, err := pkgBundle.Graph.ModuleInitialBlock(outputModule)
		if err != nil {
			return fmt.Errorf("getting module start block: %w", err)
		}
		startBlock = int64(sb)
	}

	authToken, authType := tools.GetAuth(cmd, "substreams-api-key-envvar", "substreams-api-token-envvar")
	substreamsClientConfig := client.NewSubstreamsClientConfig(
		endpoint,
		authToken,
		authType,
		sflags.MustGetBool(cmd, "insecure"),
		sflags.MustGetBool(cmd, "plaintext"),
	)

	ssClient, connClose, callOpts, headers, err := client.NewSubstreamsClient(substreamsClientConfig)
	if err != nil {
		return fmt.Errorf("substreams client setup: %w", err)
	}
	defer connClose()

	cursorStr := sflags.MustGetString(cmd, "cursor")

	stopBlock, err := readStopBlockFlag(cmd, startBlock, "stop-block", cursorStr != "")
	if err != nil {
		return fmt.Errorf("stop block: %w", err)
	}

	req := &pbsubstreamsrpc.Request{
		StartBlockNum:                       startBlock,
		StartCursor:                         cursorStr,
		StopBlockNum:                        stopBlock,
		FinalBlocksOnly:                     sflags.MustGetBool(cmd, "final-blocks-only"),
		Modules:                             pkgBundle.Package.Modules,
		OutputModule:                        outputModule,
		ProductionMode:                      productionMode,
		DebugInitialStoreSnapshotForModules: debugModulesInitialSnapshot,
	}

	if err := req.Validate(); err != nil {
		return fmt.Errorf("validate request: %w", err)
	}
	toPrint := debugModulesOutput
	if toPrint == nil {
		toPrint = []string{outputModule}
	}

	ui := tui.New(req, pkgBundle.Package, toPrint)
	if err := ui.Init(outputMode); err != nil {
		return fmt.Errorf("TUI initialization: %w", err)
	}
	defer ui.CleanUpTerminal()

	streamCtx, cancel := context.WithCancel(ctx)
	ui.OnTerminated(func(err error) {
		if err != nil {
			fmt.Printf("UI terminated with error %q\n", err)
		}

		cancel()
	})
	defer cancel()

	// add additional authorization headers
	if headers.IsSet() {
		streamCtx = metadata.AppendToOutgoingContext(streamCtx, headers.ToArray()...)
	}
	//parse additional-headers flag
	additionalHeaders := sflags.MustGetStringSlice(cmd, "header")
	if additionalHeaders != nil {
		res := parseHeaders(additionalHeaders)
		headerArray := make([]string, 0, len(res)*2)
		for k, v := range res {
			headerArray = append(headerArray, k, v)
		}
		streamCtx = metadata.AppendToOutgoingContext(streamCtx, headerArray...)
	}

	ui.SetRequest(req)
	ui.Connecting()
	cli, err := ssClient.Blocks(streamCtx, req, callOpts...)
	if err != nil && streamCtx.Err() != context.Canceled {
		return fmt.Errorf("call sf.substreams.rpc.v2.Stream/Blocks: %w", err)
	}
	ui.Connected()

	for {
		resp, err := cli.Recv()
		if resp != nil {
			if err := ui.IncomingMessage(ctx, resp, testRunner); err != nil {
				fmt.Printf("RETURN HANDLER ERROR: %s\n", err)
			}
		}
		if err != nil {
			if err == io.EOF {
				ui.Cancel()
				fmt.Println("Total Read Bytes (server-side consumption):", ui.TotalReadBytes)
				fmt.Println("all done")
				if testRunner != nil {
					testRunner.LogResults()
				}

				return nil
			}

			// Special handling if interrupted the context ourselves, no error
			if streamCtx.Err() == context.Canceled {
				ui.Cancel()
				return nil
			}

			return err
		}
	}
}
