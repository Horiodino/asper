package qemu

import (
	"log"

	"github.com/horiodino/asper/internal/pkg/network"
	"github.com/horiodino/asper/internal/pkg/vm"
	"github.com/horiodino/asper/internal/utils"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

func (q *QEMU) Validate(value string) (xmlstring []byte, err error) {
	switch value {
	case utils.NETWORK:
		xmlstring, err = validateNetworkInterface(q.ResourceType.NetworkInterfaceParams)
	case utils.VIRTUALMACHINE:
		xmlstring, err = validateVirtualMachine(q.ResourceType.RunInstanceInput)
	case utils.VIRTUALMACHINEDISK:
		xmlstring, err = validateVirtualMachineDisk(q.ResourceType.InputDiskConfiguration)
	case utils.FIREWALL:
		xmlstring, err = validateFirewall(q.ResourceType.FirewallConfiguration)
	default:
	}

	return xmlstring, err

}

func validateNetworkInterface(input *network.NetworkInterfaceParams) (xmlstring []byte, err error) {

	str := libvirtxml.Network{
		IPv6:                "yes",
		TrustGuestRxFilters: "yes",
		Name:                input.Name,
		Metadata: &libvirtxml.NetworkMetadata{
			XML: "value",
		},
		Bridge: &libvirtxml.NetworkBridge{
			Name:            "bidge-adapter",
			STP:             "on",
			Delay:           "0",
			MACTableManager: "libvirt",
			Zone:            "trusted",
		},
		MTU: &libvirtxml.NetworkMTU{
			Size: uint(1500),
		},
		MAC: &libvirtxml.NetworkMAC{
			Address: "52:54:00:00:00:00",
		},
		Domain: &libvirtxml.NetworkDomain{
			Name:      "example.com",
			LocalOnly: "yes",
		},
		DNS:         &libvirtxml.NetworkDNS{},
		Bandwidth:   &libvirtxml.NetworkBandwidth{},
		PortOptions: &libvirtxml.NetworkPortOptions{},
		PortGroups:  []libvirtxml.NetworkPortGroup{},
	}

	xmlstring, err = str.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	return xmlstring, err
}

func validateVirtualMachine(input *vm.RunInstanceInput) (xmlstring []byte, err error) {

	return xmlstring, err
}

func validateVirtualMachineDisk(input *vm.InputDiskConfiguration) (xmlstring []byte, err error) {

	return xmlstring, err
}

func validateFirewall(input *vm.FirewallConfiguration) (xmlstring []byte, err error) {

	return xmlstring, err
}
