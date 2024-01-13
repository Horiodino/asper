package vm

import (
	"context"

	"github.com/horiodino/asper/internal/logger"
	// "github.com/horiodino/asper/bkend/hypervisor/qemu"
)

// TODO add use valid  argument for all methods

type vm struct {
	Token string `json:"token"` // Authentication token for the VM operations
}

func InitializeFromConfig() *vm {
	return &vm{}
}

func (c *vm) CreateVM(ctx context.Context, input InstanceConfigurationInput) (*InstanceConfigurationOutput, error) {

	response, err := validateInput(input)
	if err != nil {
		return nil, err
	}
	logger := logger.NewLogger()
	logger.LogInfo(response, input)

	return &InstanceConfigurationOutput{}, nil
}

func (c *vm) InstanceState(ctx context.Context, input DescribeInstancesInput) (*DescribeInstancesOutput, error) {
	return &DescribeInstancesOutput{}, nil
}

func (c *vm) ModifyVM(ctx context.Context, input InstanceConfigurationInput) (*InstanceConfigurationOutput, error) {
	return &InstanceConfigurationOutput{}, nil
}

func (c *vm) MoveVM(ctx context.Context, input InstanceConfigurationInput) (*InstanceConfigurationOutput, error) {
	return &InstanceConfigurationOutput{}, nil
}

func (c *vm) ResizeVM(ctx context.Context, input InstanceConfigurationInput) (*InstanceConfigurationOutput, error) {
	return &InstanceConfigurationOutput{}, nil
}

func (c *vm) RebootVM(ctx context.Context, input DescribeInstancesInput) (*DescribeInstancesOutput, error) {
	return &DescribeInstancesOutput{}, nil
}

func (c *vm) StopVM(ctx context.Context, input DescribeInstancesInput) (*DescribeInstancesOutput, error) {
	return &DescribeInstancesOutput{}, nil
}

func (c *vm) StartVM(ctx context.Context, input DescribeInstancesInput) (*DescribeInstancesOutput, error) {
	return &DescribeInstancesOutput{}, nil
}

func (c *vm) RebuildVM(ctx context.Context, input InstanceConfigurationInput) (*InstanceConfigurationOutput, error) {
	return &InstanceConfigurationOutput{}, nil
}

func (c *vm) ReinstallVM(ctx context.Context, input InstanceConfigurationInput) (*InstanceConfigurationOutput, error) {
	return &InstanceConfigurationOutput{}, nil
}

func (c *vm) ResetVM(ctx context.Context, input DescribeInstancesInput) (*DescribeInstancesOutput, error) {
	return &DescribeInstancesOutput{}, nil
}

func (c *vm) CloneVM(ctx context.Context, input InstanceConfigurationInput) (*InstanceConfigurationOutput, error) {
	return &InstanceConfigurationOutput{}, nil
}

func (c *vm) BackupVM(ctx context.Context, input InstanceConfigurationInput) (*InstanceConfigurationOutput, error) {
	return &InstanceConfigurationOutput{}, nil
}

func (c *vm) RestoreVM(ctx context.Context, input InstanceConfigurationInput) (*InstanceConfigurationOutput, error) {
	return &InstanceConfigurationOutput{}, nil
}

func (c *vm) ConvertVM(ctx context.Context, input InstanceConfigurationInput) (*InstanceConfigurationOutput, error) {
	return &InstanceConfigurationOutput{}, nil
}
