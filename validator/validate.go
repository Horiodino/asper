package validator

import (
	"context"

	"github.com/horiodino/asper/internal/configuration"
	"github.com/horiodino/asper/internal/logger"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

const (
	DiskConfigurationInput     = "DiskConfigurationInput"
	NetworkConfigurationInput  = "NetworkConfigurationInput"
	FirewallConfigurationInput = "FirewallConfigurationInput"
	SSHConfigurationInput      = "SSHConfigurationInput"
	InstanceConfigurationInput = "InstanceConfigurationInput"
)

type Validator interface {
	IsValidfields(cxt context.Context, input interface{}) (bool, error)
}
type ValidatorClient struct {
	// it will cointain  the * to the orignal configration.InputDiskConfiguration and all the other thins user defined
	config *configuration.ALL
}

func InitializeValidator(c *configuration.ALL) *ValidatorClient {
	return &ValidatorClient{
		config: c,
	}
}

type InstanceConfigurationOutput struct {
}

func (c *ValidatorClient) IsValidfields(cxt context.Context, input interface{}) (string, error) {
	switch input {
	case DiskConfigurationInput:
		return "", nil
	case NetworkConfigurationInput:
		return "", nil
	case FirewallConfigurationInput:
		return "", nil
	case SSHConfigurationInput:
		return "", nil
	case InstanceConfigurationInput:
		logger := logger.NewLogger()
		logger.LogInfo("isVAlidfiedsRunning!!")
		resource := c.config.InstanceConfigurationInput
		config, err := isValidInstanceConfiguration(cxt, resource)
		if err != nil {
			logger.LogInfo("Error in isValidInstanceConfiguration")
			return "", err
		}

		logger.LogInfo("All fields are valid")
		return config, err
	default:
		return "", nil
	}
}

func isValidInstanceConfiguration(cxt context.Context, input configuration.InstanceConfigurationInput) (string, error) {

	var path string

	switch input.OS {
	case "ubuntu":
		path = "/var/lib/libvirt/images/ubuntu-20.04-server-cloudimg-amd64.img"
	case "centos":
		path = "/var/lib/libvirt/images/CentOS-7-x86_64-GenericCloud.qcow2"
	}

	domain := &libvirtxml.Domain{

		Type: "qemu",
		Name: string(input.Name),
		Memory: &libvirtxml.DomainMemory{
			Unit:  "KiB",
			Value: input.Memory, // 1 GiB
		},

		MaximumMemory: &libvirtxml.DomainMaxMemory{
			Unit:  "KiB",
			Value: input.Memory, // 2 GiB
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
							File: path,
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
		OnPoweroff: "destroy",
		OnReboot:   "restart",
		OnCrash:    "restart",
	}
	xml, _ := domain.Marshal()
	return xml, nil
}
