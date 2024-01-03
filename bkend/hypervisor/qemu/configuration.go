package qemu

import (
	"fmt"
	"log"

	"github.com/horiodino/asper/internal/pkg/network"
	"github.com/horiodino/asper/internal/pkg/vm"
	libvertgo "github.com/libvirt/libvirt-go"
)

type QEMU struct {
	ResourceName string
	ResourceType *all
	Client       *libvertgo.Connect
}

type all struct {
	InstanceConfigurationInput   *vm.InstanceConfigurationInput
	InstanceConfigurationOutput  *vm.InstanceConfigurationOutput
	InputDiskConfiguration       *vm.InputDiskConfiguration
	InputNetworkInterface        *vm.InputNetworkInterface
	InputSSHConfiguration        *vm.InputSSHConfiguration
	RunInstanceInput             *vm.RunInstanceInput
	RunInstanceOutput            *vm.RunInstanceOutput
	FirewallConfiguration        *vm.FirewallConfiguration
	FirewallRule                 *vm.FirewallRule
	InboundRule                  *vm.InboundRule
	OutboundRule                 *vm.OutboundRule
	NetworkInterfaceParams       *network.NetworkInterfaceParams
	InterfaceResult              *network.InterfaceResult
	DeleteNetworkInterfaceParams *network.DeleteNetworkInterfaceParams
	DeleteNetworkInterfaceResult *network.DeleteNetworkInterfaceResult
	AttachNetworkInterfaceParams *network.AttachNetworkInterfaceParams
	AttachNetworkInterfaceResult *network.AttachNetworkInterfaceResult
	DetachNetworkInterfaceParams *network.DetachNetworkInterfaceParams
	DetachNetworkInterfaceResult *network.DetachNetworkInterfaceResult
	BridgeConfigurationInput     *network.BridgeConfigurationInput
}

func qemuclient() *libvertgo.Connect {
	conn, err := libvertgo.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("Error connecting to hypervisor:", err)
		log.Fatal(err)
	}

	return conn
}

func InitializeQEMU() *QEMU {
	return &QEMU{
		ResourceName: "nil",
		ResourceType: "nil",
		Client:       qemuclient(),
	}
}
