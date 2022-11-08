package outputgraph

import (
	"fmt"
	"github.com/streamingfast/substreams/manifest"
	pbsubstreams "github.com/streamingfast/substreams/pb/sf/substreams/v1"
)

func ValidateRequest(request *pbsubstreams.Request, blockType string) error {
	if request.StartBlockNum < 0 {
		// TODO(abourget): remove this check once we support StartBlockNum being negative
		return fmt.Errorf("negative start block %d is not accepted", request.StartBlockNum)
	}

	if request.Modules == nil {
		return fmt.Errorf("no modules found in request")
	}

	if err := manifest.ValidateModules(request.Modules); err != nil {
		return fmt.Errorf("modules validation failed: %w", err)
	}

	if err := pbsubstreams.ValidateRequest(request); err != nil {
		return fmt.Errorf("validate request: %s", err)
	}

	for _, binary := range request.Modules.Binaries {
		if binary.Type != "wasm/rust-v1" {
			return fmt.Errorf(`unsupported binary type: %q, please use "wasm/rust-v1"`, binary.Type)
		}
	}

	for _, mod := range request.Modules.Modules {
		for _, input := range mod.Inputs {
			if src := input.GetSource(); src != nil {
				if src.Type != blockType && src.Type != "sf.substreams.v1.Clock" {
					return fmt.Errorf("input source %q not supported, only %q and 'sf.substreams.v1.Clock' are valid", src, blockType)
				}
			}
		}
	}

	return nil
}