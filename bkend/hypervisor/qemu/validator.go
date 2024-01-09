package qemu

import (
	"fmt"
	"log"

	"github.com/horiodino/asper/internal/utils"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

func (q *QEMU) Validate(value string) (xmlstring string, err error) {
	switch value {
	case utils.NETWORK:
		xmlstring, err = q.validateNetworkInterface()
	case utils.VIRTUALMACHINE:
		xmlstring, err = q.validateVirtualMachine()
	//	case utils.VIRTUALMACHINEDISK:
	//		xmlstring, err = validateVirtualMachineDisk(q.ResourceType.InputDiskConfiguration)
	//	case utils.FIREWALL:
	//		xmlstring, err = validateFirewall(q.ResourceType.FirewallConfiguration)
	default:
		fmt.Println("which opration ??")
	}

	return xmlstring, err

}

func (q *QEMU) validateNetworkInterface() (xmlstring string, err error) {

	str := libvirtxml.Network{
		IPv6:                "yes",
		TrustGuestRxFilters: "yes",
		Name:                "asper-1-nic",
		Bridge: &libvirtxml.NetworkBridge{
			Name:            "asper-1-bridge",
			STP:             "on",
			Delay:           "0",
			MACTableManager: "libvirt",
			Zone:            "trusted",
		},
		Forward: &libvirtxml.NetworkForward{
			Mode: "nat",
		},
		MTU: &libvirtxml.NetworkMTU{
			Size: uint(1500),
		},
		//		MAC: &libvirtxml.NetworkMAC{
		//			Address: "52:54:00:00:00:00",
		//		},
		DNS:         &libvirtxml.NetworkDNS{},
		Bandwidth:   &libvirtxml.NetworkBandwidth{},
		PortOptions: &libvirtxml.NetworkPortOptions{},
		IPs: []libvirtxml.NetworkIP{
			{
				Address: "192.168.1.2",
				Family:  "ipv4",
				Prefix:  24,
				DHCP: &libvirtxml.NetworkDHCP{
					Ranges: []libvirtxml.NetworkDHCPRange{
						{
							Start: "192.168.1.50",  // Adjust the start of the range within the network subnet
							End:   "192.168.1.254", // Adjust the end of the range within the network subnet
						},
					},
				},
			},
		},
	}

	xmlstring, err = str.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	return xmlstring, err
}

func (q *QEMU) validateVirtualMachine() (xmlstring string, err error) {

	domain := &libvirtxml.Domain{
		Type: "qemu",
		Name: "ab",
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
							File: "/var/lib/libvirt/images/ubuntu.img",
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

			// Interfaces: []libvirtxml.DomainInterface{
			// 	{
			// 		//					MAC: &libvirtxml.DomainInterfaceMAC{
			// 		//						Address: "52:54:00:01:23:45",
			// 		//					},
			// 		//					Model: &libvirtxml.DomainInterfaceModel{
			// 		//						Type: "virtio",
			// 		//					},
			// 		Source: &libvirtxml.DomainInterfaceSource{
			// 			Network: &libvirtxml.DomainInterfaceSourceNetwork{
			// 				// Network: "enp0s3",
			// 				Bridge: "br10",
			// 			},
			// 		},
			// 	},
			// },
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
		// OnPoweroff: "destroy",
		// OnReboot:   "restart",
		// OnCrash:    "restart",
	}

	xmlstring, err = domain.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	return xmlstring, err
}

//func validateVirtualMachineDisk(input *vm.InputDiskConfiguration) (xmlstring []byte, err error) {
//
//	return xmlstring, err
//}
//
//func validateFirewall(input *vm.FirewallConfiguration) (xmlstring []byte, err error) {
//
//	return xmlstring, err
//}
