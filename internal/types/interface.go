package types

import (
	"context"
	"github.com/horiodino/asper/internal/pkg/vm"
)

type VirtualLization struct {
	HyperVisor Instance
}

// Instance interface groups together methods for creating various components of a VM
type Instance interface {
	CreateVM(ctx context.Context, input vm.InstanceConfigurationInput) (vm.InstanceConfigurationOutput, error)
	DeleteVM(ctx context.Context, input vm.DescribeInstancesInput) (vm.DescribeInstancesOutput, error)
	InstanceState(context.Context, vm.DescribeInstancesInput) (*vm.DescribeInstancesOutput, error)
	ModifyVM(ctx context.Context, input vm.InstanceConfigurationInput) (*vm.InstanceConfigurationOutput, error)
	MoveVM(ctx context.Context, input vm.InstanceConfigurationInput) (*vm.InstanceConfigurationOutput, error)
	ResizeVM(ctx context.Context, input vm.InstanceConfigurationInput) (*vm.InstanceConfigurationOutput, error)
	RebootVM(ctx context.Context, input vm.DescribeInstancesInput) (*vm.DescribeInstancesOutput, error)
	StopVM(ctx context.Context, input vm.DescribeInstancesInput) (*vm.DescribeInstancesOutput, error)
	StartVM(ctx context.Context, input vm.DescribeInstancesInput) (vm.*DescribeInstancesOutput, error)
	RebuildVM(ctx context.Context, input vm.InstanceConfigurationInput) (*vm.InstanceConfigurationOutput, error)
	ReinstallVM(ctx context.Context, input vm.InstanceConfigurationInput) (*vm.InstanceConfigurationOutput, error)
	ResetVM(ctx context.Context, input vm.DescribeInstancesInput) (*vm.DescribeInstancesOutput, error)
	CloneVM(ctx context.Context, input vm.InstanceConfigurationInput) (*vm.InstanceConfigurationOutput, error)
	BackupVM(ctx context.Context, input vm.InstanceConfigurationInput) (*vm.InstanceConfigurationOutput, error)
	RestoreVM(ctx context.Context, input vm.InstanceConfigurationInput) (*vm.InstanceConfigurationOutput, error)
	ConvertVM(ctx context.Context, input vm.InstanceConfigurationInput) (*vm.InstanceConfigurationOutput, error)
	DescribeInstances(ctx context.Context, input vm.DescribeInstancesInput) (*vm.DescribeInstancesOutput, error)
}
