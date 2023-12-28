package main

import (
	"fmt"

	libvirt "github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

func main() {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("Error connecting to hypervisor:", err)
		return
	}
	defer conn.Close()

	domainDef := createDomainDef()

	domain, err := conn.DomainDefineXML(domainDef)
	if err != nil {
		fmt.Println("Error defining domain:", err)
		return
	}
	defer domain.Free()

	err = domain.Create()
	if err != nil {
		fmt.Println("Error starting domain:", err)
		return
	}

	fmt.Println("Virtual machine started successfully.")
}

func createDomainDef() string {

	domain := &libvirtxml.Domain{

		Type: "qemu",
		Name: "jatatadharishiv",
		Memory: &libvirtxml.DomainMemory{
			Unit:  "KiB",
			Value: 1048576, // 1 GiB
		},

		MaximumMemory: &libvirtxml.DomainMaxMemory{
			Unit:  "KiB",
			Value: 1097152, // 2 GiB
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
	return xml
}
