package exec

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/streamingfast/substreams/metrics"
	pbssinternal "github.com/streamingfast/substreams/pb/sf/substreams/intern/v2"
	pbsubstreams "github.com/streamingfast/substreams/pb/sf/substreams/v1"
	"github.com/streamingfast/substreams/reqctx"
	"github.com/streamingfast/substreams/storage/execout"
	"github.com/streamingfast/substreams/storage/index"
)

type MockExecOutput struct {
	clockFunc func() *pbsubstreams.Clock

	cacheMap map[string][]byte
}

func (t *MockExecOutput) Clock() *pbsubstreams.Clock {
	return t.clockFunc()
}

func (t *MockExecOutput) Len() int {
	return 0
}

func (t *MockExecOutput) Get(name string) ([]byte, bool, error) {
	v, ok := t.cacheMap[name]
	if !ok {
		return nil, false, execout.ErrNotFound
	}
	return v, true, nil
}

func (t *MockExecOutput) Set(name string, value []byte) (err error) {
	t.cacheMap[name] = value
	return nil
}

type MockModuleExecutor struct {
	name string

	RunFunc      func(ctx context.Context, reader execout.ExecutionOutputGetter) (out []byte, outForFiles []byte, moduleOutputData *pbssinternal.ModuleOutput, err error)
	ApplyFunc    func(value []byte) error
	LogsFunc     func() (logs []string, truncated bool)
	StackFunc    func() []string
	ToOutputFunc func(data []byte) (*pbssinternal.ModuleOutput, error)
	cacheable    bool
}

var _ ModuleExecutor = (*MockModuleExecutor)(nil)

func (t *MockModuleExecutor) run(ctx context.Context, reader execout.ExecutionOutputGetter) (out []byte, outForFiles []byte, moduleOutputData *pbssinternal.ModuleOutput, err error) {
	if t.RunFunc != nil {
		return t.RunFunc(ctx, reader)
	}
	return nil, nil, nil, fmt.Errorf("not implemented")
}
func (t *MockModuleExecutor) BlockIndex() *index.BlockIndex   { return nil }
func (t *MockModuleExecutor) RunsOnBlock(_ uint64) bool       { return true }
func (t *MockModuleExecutor) Name() string                    { return t.name }
func (t *MockModuleExecutor) String() string                  { return fmt.Sprintf("TestModuleExecutor(%s)", t.name) }
func (t *MockModuleExecutor) Close(ctx context.Context) error { return nil }
func (t *MockModuleExecutor) HasValidOutput() bool            { return t.cacheable }
func (t *MockModuleExecutor) HasOutputForFiles() bool         { return false }

func (t *MockModuleExecutor) applyCachedOutput(value []byte) error {
	if t.ApplyFunc != nil {
		return t.ApplyFunc(value)
	}
	return fmt.Errorf("not implemented")
}

func (t *MockModuleExecutor) toModuleOutput(data []byte) (*pbssinternal.ModuleOutput, error) {
	if t.ToOutputFunc != nil {
		return t.ToOutputFunc(data)
	}
	return nil, fmt.Errorf("not implemented")
}

func (t *MockModuleExecutor) lastExecutionLogs() (logs []string, truncated bool) {
	if t.LogsFunc != nil {
		return t.LogsFunc()
	}
	return nil, false
}

func TestModuleExecutorRunner_Run_HappyPath(t *testing.T) {
	ctx := context.Background()

	ctx = reqctx.WithReqStats(ctx, metrics.NewReqStats(&metrics.Config{}, zap.NewNop()))
	executor := &MockModuleExecutor{
		name: "test",
		RunFunc: func(ctx context.Context, reader execout.ExecutionOutputGetter) (out []byte, outForFiles []byte, moduleOutputData *pbssinternal.ModuleOutput, err error) {
			return []byte("test"), nil, &pbssinternal.ModuleOutput{
				Data: &pbssinternal.ModuleOutput_MapOutput{
					MapOutput: nil,
				},
			}, nil
		},
		LogsFunc: func() (logs []string, truncated bool) {
			return []string{"test"}, false
		},
	}
	output := &MockExecOutput{
		cacheMap: make(map[string][]byte),
	}

	moduleOutput, _, _, _, err := RunModule(ctx, executor, output)
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, moduleOutput)
}

func TestModuleExecutorRunner_Run_CachedOutput(t *testing.T) {
	ctx := context.Background()

	applied := false

	executor := &MockModuleExecutor{
		name: "test",
		RunFunc: func(ctx context.Context, reader execout.ExecutionOutputGetter) (out []byte, outForFiles []byte, moduleOutputData *pbssinternal.ModuleOutput, err error) {
			return []byte("test"), nil, &pbssinternal.ModuleOutput{
				Data: &pbssinternal.ModuleOutput_MapOutput{
					MapOutput: nil,
				},
			}, nil
		},
		ToOutputFunc: func(data []byte) (*pbssinternal.ModuleOutput, error) {
			return &pbssinternal.ModuleOutput{
				Data: &pbssinternal.ModuleOutput_MapOutput{
					MapOutput: nil,
				},
			}, nil
		},
		ApplyFunc: func(value []byte) error {
			applied = true
			return nil
		},
		LogsFunc: func() (logs []string, truncated bool) {
			return []string{"test"}, false
		},
	}
	output := &MockExecOutput{
		cacheMap: map[string][]byte{
			"test": []byte("cached"),
		},
	}

	moduleOutput, _, _, _, err := RunModule(ctx, executor, output)
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
	assert.True(t, applied)
	assert.NotEmpty(t, moduleOutput)
	assert.True(t, moduleOutput.Cached)
}
