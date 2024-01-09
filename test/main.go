package main

import (
	"context"
	"log"

	"github.com/horiodino/asper/bkend/hypervisor/qemu"
	"github.com/horiodino/asper/internal/pkg/vm"
	// "github.com/horiodino/asper/internal/pkg/network"
)

func main() {
	ctx := context.Background()
	//	Nclient := network.NewNetworkClient()
	//	_, err := Nclient.CreateNetworkInterface(ctx,&qemu.NetworkInterfaceParams{
	//		Name: "testing-nic",
	//	})
	//	if err != nil {
	//		log.Fatal(err)
	//	}

	Client := vm.InitializeFromConfig()
	_, err := Client.CreateVM(ctx, &qemu.InstanceConfigurationInput{
		Name: "demo-asper-vm",
	})
	if err != nil {
		log.Fatal(err)
	}

}

// package main

// import (
// 	"fmt"

// 	"github.com/libvirt/libvirt-go"
// 	libvirtxml "github.com/libvirt/libvirt-go-xml"
// )

// func main() {
// 	// Connect to the local QEMU instance using the URI
// 	conn, err := libvirt.NewConnect("qemu:///system")
// 	if err != nil {
// 		fmt.Println("Failed to connect to QEMU:", err)
// 		return
// 	}
// 	defer conn.Close()

// 	// Create a new Network definition using libvirtxml
// 	networkDef := &libvirtxml.Network{
// 		Name: "virtual-net",
// 		Bridge: &libvirtxml.NetworkBridge{
// 			Name: "virtual-bridge01",
// 			STP:  "on",
// 		},
// 		Forward: &libvirtxml.NetworkForward{
// 			Mode: "nat",
// 		},
// 		IPs: []libvirtxml.NetworkIP{
// 			{
// 				Address: "192.168.100.1",
// 				Netmask: "255.255.255.0",
// 				DHCP:    &libvirtxml.NetworkDHCP{Ranges: []libvirtxml.NetworkDHCPRange{{Start: "192.168.100.128", End: "192.168.100.254"}}},
// 			},
// 		},
// 		MTU: &libvirtxml.NetworkMTU{
// 			Size: 1500,
// 		},
// 		MAC: &libvirtxml.NetworkMAC{
// 			Address: "52:54:00:00:00:01",
// 		},
// 	}

// 	// Define the network using the parsed definition
// 	xmlData, err := networkDef.Marshal()
// 	if err != nil {
// 		fmt.Println("Failed to marshal network definition:", err)
// 		return
// 	}

// 	network, err := conn.NetworkDefineXML(string(xmlData))
// 	if err != nil {
// 		fmt.Println("Failed to define network:", err)
// 		return
// 	}

// 	err = network.Create()
// 	if err != nil {
// 		fmt.Println("Failed to start network:", err)
// 		return
// 	}

// 	fmt.Println("Network interface created successfully")
// }
