package vm

import (
	"context"
	"fmt"
	"log"

	"github.com/horiodino/asper/bkend/hypervisor/qemu"
)

// TODO add use valid  argument for all methods

type vm struct {
	Token string `json:"token"` // Authentication token for the VM operations
}

func InitializeFromConfig() *vm {
	return &vm{}
}

// Instance interface groups together methods for creating various components of a VM
type Instance interface {
	CreateVM(ctx context.Context, input InstanceConfigurationInput) (InstanceConfigurationOutput, error)
	DeleteVM(ctx context.Context, input DescribeInstancesInput) (DescribeInstancesOutput, error)
	InstanceState(context.Context, DescribeInstancesInput) (*DescribeInstancesOutput, error)
	InstanceSSHKey(ctx context.Context, input DescribeSSHKeyInput) (*DescribeSSHKeyOutput, error)
	InstanceNetworkInterface(ctx context.Context, input DescribeNetworkInterfaceInput) (*DescribeNetworkInterfaceOutput, error)
	InstanceDiskConfiguration(ctx context.Context, input DescribeDiskConfigurationInput) (*DescribeDiskConfigurationOutput, error)
	InstanceFirewallConfiguration(ctx context.Context, input DescribeFirewallConfigurationInput) (*DescribeFirewallConfigurationOutput, error)
	ModifyVM(ctx context.Context, input InstanceConfigurationInput) (*InstanceConfigurationOutput, error)
	MoveVM(ctx context.Context, input InstanceConfigurationInput) (*InstanceConfigurationOutput, error)
	ResizeVM(ctx context.Context, input InstanceConfigurationInput) (*InstanceConfigurationOutput, error)
	RebootVM(ctx context.Context, input DescribeInstancesInput) (*DescribeInstancesOutput, error)
	StopVM(ctx context.Context, input DescribeInstancesInput) (*DescribeInstancesOutput, error)
	StartVM(ctx context.Context, input DescribeInstancesInput) (*DescribeInstancesOutput, error)
	RebuildVM(ctx context.Context, input InstanceConfigurationInput) (*InstanceConfigurationOutput, error)
	ReinstallVM(ctx context.Context, input InstanceConfigurationInput) (*InstanceConfigurationOutput, error)
	ResetVM(ctx context.Context, input DescribeInstancesInput) (*DescribeInstancesOutput, error)
	CloneVM(ctx context.Context, input InstanceConfigurationInput) (*InstanceConfigurationOutput, error)
	BackupVM(ctx context.Context, input InstanceConfigurationInput) (*InstanceConfigurationOutput, error)
	RestoreVM(ctx context.Context, input InstanceConfigurationInput) (*InstanceConfigurationOutput, error)
	ConvertVM(ctx context.Context, input InstanceConfigurationInput) (*InstanceConfigurationOutput, error)
	DescribeInstances(ctx context.Context, input DescribeInstancesInput) (*DescribeInstancesOutput, error)
	DescribeSSHKey(ctx context.Context, input DescribeSSHKeyInput) (*DescribeSSHKeyOutput, error)
	DescribeNetworkInterface(ctx context.Context, input DescribeNetworkInterfaceInput) (*DescribeNetworkInterfaceOutput, error)
	DescribeDiskConfiguration(ctx context.Context, input DescribeDiskConfigurationInput) (*DescribeDiskConfigurationOutput, error)
	DescribeFirewallConfiguration(ctx context.Context, input DescribeFirewallConfigurationInput) (*DescribeFirewallConfigurationOutput, error)
}

func (c *vm) CreateVM(ctx context.Context, input InstanceConfigurationInput) (*InstanceConfigurationOutput, error) {

	QEMU := qemu.InitializeQEMU()
	xmlvalue, err := QEMU.Validate("vm")
	if err != nil {
		log.Fatal(err)
	}

	val, err := QEMU.Client.DomainDefineXML(xmlvalue)
	if err != nil {
		log.Fatal(err)
	}

	val.Create()
	fmt.Println(val.GetInfo())

	return &InstanceConfigurationOutput{}, nil
}

func (c *vm) InstanceState(ctx context.Context, input DescribeInstancesInput) (*DescribeInstancesOutput, error) {
	return &DescribeInstancesOutput{}, nil
}

func (c *vm) InstanceSSHKey(ctx context.Context, input DescribeSSHKeyInput) (*DescribeSSHKeyOutput, error) {
	return &DescribeSSHKeyOutput{}, nil
}

func (c *vm) InstanceNetworkInterface(ctx context.Context, input DescribeNetworkInterfaceInput) (*DescribeNetworkInterfaceOutput, error) {
	return &DescribeNetworkInterfaceOutput{}, nil
}

func (c *vm) InstanceDiskConfiguration(ctx context.Context, input DescribeDiskConfigurationInput) (*DescribeDiskConfigurationOutput, error) {
	return &DescribeDiskConfigurationOutput{}, nil
}

func (c *vm) InstanceFirewallConfiguration(ctx context.Context, input DescribeFirewallConfigurationInput) (*DescribeFirewallConfigurationOutput, error) {
	return &DescribeFirewallConfigurationOutput{}, nil
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
