package network

import (
	"context"
	"fmt"
	"log"

	"github.com/horiodino/asper/bkend/hypervisor/qemu"
)

// Network the newtwork package allocates networking resources required
type Network interface {
	CreateNetworkInterface(ctx context.Context, nicinput *NetworkInterfaceParams) (result InterfaceResult, err error)
	DeleteNetworkInterface(ctx context.Context, nicinput *DeleteNetworkInterfaceParams) (result DeleteNetworkInterfaceResult, err error)
	AttachNetworkInterface(ctx context.Context, nicinput *AttachNetworkInterfaceParams) (result AttachNetworkInterfaceResult, err error)
	DetachNetworkInterface(ctx context.Context, nicinput *DetachNetworkInterfaceParams) (result DetachNetworkInterfaceResult, err error)
	CreateFirewall(ctx context.Context, nicinput *NetworkInterfaceParams) (result InterfaceResult, err error)
}

type network struct {
	Token string `json:"token"`
}

func NewNetworkClient() Network {
	return &network{}
}

// CreateFirewall creates a firewall
func (n *network) CreateFirewall(ctx context.Context, nicinput *NetworkInterfaceParams) (result InterfaceResult, err error) {
	return result, err
}

// CreateNetworkInterface creates a new network interface
func (n *network) CreateNetworkInterface(ctx context.Context, nicinput *NetworkInterfaceParams) (result InterfaceResult, err error) {

	QEMU := qemu.InitializeQEMU()
	xmlstring, err := QEMU.Validate("network")
	nic, err := QEMU.Client.NetworkDefineXML(string(xmlstring))
	if err != nil {
		log.Fatal(err)
	}

	nic.Create()

	fmt.Println(nic.GetName())
	fmt.Println(nic.GetBridgeName())
	fmt.Println(nic.IsActive())

	ctx.Done()
	return result, err
}

// DeleteNetworkInterface deletes a network interface
func (n *network) DeleteNetworkInterface(ctx context.Context, nicinput *DeleteNetworkInterfaceParams) (result DeleteNetworkInterfaceResult, err error) {
	return result, err
}

// AttachNetworkInterface attaches a network interface
func (n *network) AttachNetworkInterface(ctx context.Context, nicinput *AttachNetworkInterfaceParams) (result AttachNetworkInterfaceResult, err error) {
	return result, err
}

// DetachNetworkInterface detaches a network interface
func (n *network) DetachNetworkInterface(ctx context.Context, nicinput *DetachNetworkInterfaceParams) (result DetachNetworkInterfaceResult, err error) {
	return result, err
}
