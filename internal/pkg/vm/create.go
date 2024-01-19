package vm

import (
	"context"

	"github.com/horiodino/asper/internal/logger"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
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

	data, err := modifyinstanceconfig(input)
	if err != nil {
		return nil, err
	}

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

func modifyinstanceconfig(input InstanceConfigurationInput) (string, error) {

	domain := &libvirtxml.Domain{

		Type: "qemu",
		Name: input.Name,
		Memory: &libvirtxml.DomainMemory{
			Unit:  "KiB",
			Value: input.Memory,
		},

		MaximumMemory: &libvirtxml.DomainMaxMemory{
			Unit:  "KiB",
			Value: input.Memory,
		},
		VCPU: &libvirtxml.DomainVCPU{
			Placement: "static",
			Value:     input.VCPU,
		},
		Devices: &libvirtxml.DomainDeviceList{
			Disks: []libvirtxml.DomainDisk{
				{
					Device: "disk",
					Driver: &libvirtxml.DomainDiskDriver{
						Name: "qemu",
						Type: "qcow2",
					},
					Source: &libvirtxml.DomainDiskSource{
						File: &libvirtxml.DomainDiskSourceFile{
							File: "/var/lib/libvirt/images/ubuntu-20.04-server-cloudimg-amd64.img",
						},
					},
					Target: &libvirtxml.DomainDiskTarget{
						Dev: "vda",
						Bus: "virtio",
					},
					BlockIO: &libvirtxml.DomainDiskBlockIO{
						LogicalBlockSize:  512,
						PhysicalBlockSize: 512,
					},
					Boot: &libvirtxml.DomainDeviceBoot{
						Order: 1,
					},
				},
			},

			Graphics: []libvirtxml.DomainGraphic{
				{
					Spice: &libvirtxml.DomainGraphicSpice{
						AutoPort: "yes",
					},
				},
			},

			Interfaces: []libvirtxml.DomainInterface{
				{
					MAC: &libvirtxml.DomainInterfaceMAC{
						Address: "52:54:00:01:23:45",
					},
					Model: &libvirtxml.DomainInterfaceModel{
						Type: "virtio",
					},
					Source: &libvirtxml.DomainInterfaceSource{
						Network: &libvirtxml.DomainInterfaceSourceNetwork{
							Network: "default",
						},
					},
					Target: &libvirtxml.DomainInterfaceTarget{
						Dev:     "vnet0",
						Managed: "yes",
					},
				},
			},
			Controllers: []libvirtxml.DomainController{},
			Leases:      []libvirtxml.DomainLease{},
			Channels:    []libvirtxml.DomainChannel{},
			Consoles:    []libvirtxml.DomainConsole{},
			RedirDevs:   []libvirtxml.DomainRedirDev{},
		},

		Metadata: &libvirtxml.DomainMetadata{
			XML: "<domain type='kvm'/>",
		},

		OS: &libvirtxml.DomainOS{
			Type: &libvirtxml.DomainOSType{
				Arch:    "x86_64",
				Machine: "pc",
				Type:    "hvm",
			},
		},
		OnPoweroff: "restart",
		OnReboot:   "restart",
		OnCrash:    "restart",
	}

	xml, _ := domain.Marshal()
	return xml, nil
}
