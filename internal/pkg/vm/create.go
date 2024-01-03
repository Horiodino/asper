package vm

import (
	"context"
	"fmt"
	"log"

	"github.com/horiodino/asper/bkend/hypervisor/qemu"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
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

	QEMU := qemu.InitializeQEMU()
	QEMU.ResourceType.InstanceConfigurationInput = input
	domain := &libvirtxml.Domain{

		Type: "qemu",
		Name: "test-vm",
		Memory: &libvirtxml.DomainMemory{
			Unit:  "KiB",
			Value: 1048576, // 1 GiB
		},
		VCPU: &libvirtxml.DomainVCPU{
			Placement: "static",
			Value:     2,
		},
		Devices: &libvirtxml.DomainDeviceList{

			// Disks
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
							Network: "net-nic",
							Bridge:  "bidge-adapter",
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
		OnPoweroff: "destroy",
		OnReboot:   "restart",
		OnCrash:    "restart",
	}
	xmlstr, err := domain.Marshal()
	if err != nil {
		log.Fatal(err)
	}

	val, err := QEMU.Client.DomainDefineXML(xmlstr)
	if err != nil {
		log.Fatal(err)
	}

	val.Create()
	fmt.Println(val.GetInfo())

	return nil, nil
}
