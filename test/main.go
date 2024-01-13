package main

import (
	"encoding/json"
	"fmt"

	"github.com/horiodino/asper/internal/pkg/disk"
	"github.com/horiodino/asper/internal/pkg/network"
	"github.com/horiodino/asper/internal/pkg/ssh"
	"github.com/horiodino/asper/internal/pkg/vm"
)

func isEmpty(input vm.InstanceConfigurationInput) error {
	// marshal the input to json
	jsondata, err := json.Marshal(input)
	if err != nil {
		return err
	}

	// unmarshal the json to map
	var data map[string]interface{}
	err = json.Unmarshal(jsondata, &data)
	if err != nil {
		return err
	}

	// iterate over the map and check if any of the value is empty
	for key, value := range data {
		if value == nil {
			return fmt.Errorf("missing required parameter", key)
		}
	}

	return nil
}

func main() {
	// create a instance of the struct
	input := vm.InstanceConfigurationInput{
		Name:   "demo-asper-vm",
		OS:     "ubuntu",
		Memory: 1024,
		VCPU:   1,
		SSHConfiguration: ssh.SSHConfigurationInput{
			Keyname: "asper",
		},
		NetworkInterface: network.NetworkInterfaceParams{
			Name:   "asper",
			IP:     "",
			Subnet: "",
		},
		DiskConfiguration: disk.DiskConfigurationInput{
			Name: "asper",
			Size: 10,
		},
	}

	// call the isEmpty function
	err := isEmpty(input)
	if err != nil {
		fmt.Println(err)
	}

}

// package main

// import (
// 	"context"
// 	"log"

// 	"github.com/horiodino/asper/internal/pkg/ssh"
// 	"github.com/horiodino/asper/internal/pkg/vm"
// 	// "github.com/horiodino/asper/internal/pkg/network"
// )

// func main() {
// 	ctx := context.Background()
// 	Client := vm.InitializeFromConfig()
// 	_, err := Client.CreateVM(ctx, vm.InstanceConfigurationInput{
// 		Name:   "demo-asper-vm",
// 		OS:     "ubuntu",
// 		Memory: 1024,
// 		VCPU:   1,
// 		Lock:   false,
// 		SSHConfiguration: ssh.SSHConfigurationInput{
// 			Keyname: "asper",
// 		},
// 		NetworkInterface: vm.NetworkInterfaceInput{
// 			Name:   "asper",
// 			Bridge: "br0",
// 		},
// 		DiskConfiguration: vm.DiskConfigurationInput{
// 			Name: "asper",
// 			Size: 10,
// 		},
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }

// // package main

// // import (
// // 	"fmt"

// // 	"github.com/libvirt/libvirt-go"
// // 	libvirtxml "github.com/libvirt/libvirt-go-xml"
// // )

// // func main() {
// // 	// Connect to the local QEMU instance using the URI
// // 	conn, err := libvirt.NewConnect("qemu:///system")
// // 	if err != nil {
// // 		fmt.Println("Failed to connect to QEMU:", err)
// // 		return
// // 	}
// // 	defer conn.Close()

// // 	// Create a new Network definition using libvirtxml
// // 	networkDef := &libvirtxml.Network{
// // 		Name: "virtual-net",
// // 		Bridge: &libvirtxml.NetworkBridge{
// // 			Name: "virtual-bridge01",
// // 			STP:  "on",
// // 		},
// // 		Forward: &libvirtxml.NetworkForward{
// // 			Mode: "nat",
// // 		},
// // 		IPs: []libvirtxml.NetworkIP{
// // 			{
// // 				Address: "192.168.100.1",
// // 				Netmask: "255.255.255.0",
// // 				DHCP:    &libvirtxml.NetworkDHCP{Ranges: []libvirtxml.NetworkDHCPRange{{Start: "192.168.100.128", End: "192.168.100.254"}}},
// // 			},
// // 		},
// // 		MTU: &libvirtxml.NetworkMTU{
// // 			Size: 1500,
// // 		},
// // 		MAC: &libvirtxml.NetworkMAC{
// // 			Address: "52:54:00:00:00:01",
// // 		},
// // 	}

// // 	// Define the network using the parsed definition
// // 	xmlData, err := networkDef.Marshal()
// // 	if err != nil {
// // 		fmt.Println("Failed to marshal network definition:", err)
// // 		return
// // 	}

// // 	network, err := conn.NetworkDefineXML(string(xmlData))
// // 	if err != nil {
// // 		fmt.Println("Failed to define network:", err)
// // 		return
// // 	}

// // 	err = network.Create()
// // 	if err != nil {
// // 		fmt.Println("Failed to start network:", err)
// // 		return
// // 	}

// // 	fmt.Println("Network interface created successfully")
// // }
