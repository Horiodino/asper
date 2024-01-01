package vm

import (
	"context"
	"fmt"
	"time"

	"github.com/horiodino/asper/internal/logger"
	"github.com/libvirt/libvirt-go"
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
	CreateVM(ctx context.Context, input *InstanceConfigurationInput) (*InstanceConfigurationOutput, error)
	DeleteVM(ctx context.Context, input *DescribeInstancesInput) (*DescribeInstancesOutput, error)
	InstanceState(context.Context, *DescribeInstancesInput) *DescribeInstancesOutput
	InstanceSSHKey(ctx context.Context, input *DescribeSSHKeyInput) *DescribeSSHKeyOutput
	InstanceNetworkInterface(ctx context.Context, input *DescribeNetworkInterfaceInput) *DescribeNetworkInterfaceOutput
	InstanceDiskConfiguration(ctx context.Context, input *DescribeDiskConfigurationInput) *DescribeDiskConfigurationOutput
	InstanceFirewallConfiguration(ctx context.Context, input *DescribeFirewallConfigurationInput) *DescribeFirewallConfigurationOutput
	ModifyVM(ctx context.Context, input *InstanceConfigurationInput) (*InstanceConfigurationOutput, error)
	MoveVM(ctx context.Context, input *InstanceConfigurationInput) (*InstanceConfigurationOutput, error)
	ResizeVM(ctx context.Context, input *InstanceConfigurationInput) (*InstanceConfigurationOutput, error)
	RebootVM(ctx context.Context, input *DescribeInstancesInput) (*DescribeInstancesOutput, error)
	StopVM(ctx context.Context, input *DescribeInstancesInput) (*DescribeInstancesOutput, error)
	StartVM(ctx context.Context, input *DescribeInstancesInput) (*DescribeInstancesOutput, error)
	RebuildVM(ctx context.Context, input *InstanceConfigurationInput) (*InstanceConfigurationOutput, error)
	ReinstallVM(ctx context.Context, input *InstanceConfigurationInput) (*InstanceConfigurationOutput, error)
	ResetVM(ctx context.Context, input *DescribeInstancesInput) (*DescribeInstancesOutput, error)
	CloneVM(ctx context.Context, input *InstanceConfigurationInput) (*InstanceConfigurationOutput, error)
	BackupVM(ctx context.Context, input *InstanceConfigurationInput) (*InstanceConfigurationOutput, error)
	RestoreVM(ctx context.Context, input *InstanceConfigurationInput) (*InstanceConfigurationOutput, error)
	ConvertVM(ctx context.Context, input *InstanceConfigurationInput) (*InstanceConfigurationOutput, error)
	DescribeInstances(ctx context.Context, input *DescribeInstancesInput) (*DescribeInstancesOutput, error)
	DescribeSSHKey(ctx context.Context, input *DescribeSSHKeyInput) (*DescribeSSHKeyOutput, error)
	DescribeNetworkInterface(ctx context.Context, input *DescribeNetworkInterfaceInput) (*DescribeNetworkInterfaceOutput, error)
	DescribeDiskConfiguration(ctx context.Context, input *DescribeDiskConfigurationInput) (*DescribeDiskConfigurationOutput, error)
	DescribeFirewallConfiguration(ctx context.Context, input *DescribeFirewallConfigurationInput) (*DescribeFirewallConfigurationOutput, error)
}

func (c *vm) CreateVM(ctx context.Context, input *InstanceConfigurationInput) (*InstanceConfigurationOutput, error) {

	// fmt.Println("working!!!!")
	// switch operationType {
	// case "managed":
	// case "custom":
	// default:
	// }

	// Creation of VM
	logger := logger.NewLogger()
	logger.LogInfo("Running CreateVM")

	time.Sleep(15 * time.Second)
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		logger.LogError(fmt.Sprintf("Error connecting to hypervisor: %v", err))
		return nil, err
	}
	defer conn.Close()

	domain, err := conn.DomainDefineXML(input)
	if err != nil {
		logger.LogError(fmt.Sprintf("Error defining domain: %v", err))
		return nil, err
	}
	defer domain.Free()

	logger.LogInfo("Initializing virtual machine components")

	err = domain.Create()
	if err != nil {
		logger.LogError(fmt.Sprintf("Error starting domain: %v", err))
		return nil, err
	}

	logger.LogInfo("Virtual machine started successfully.")
	logger.LogInfo("Initializing virtual machine components. Please wait for a few seconds.")

	return nil, nil
}
