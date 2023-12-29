package network

import "context"

// Network the newtwork package allocates networking resources required
type Network interface {
	CreateNetworkInterface(ctx context.Context, nicinput *NetworkInterfaceParams) (result InterfaceResult, err error)
	DeleteNetworkInterface(ctx context.Context, nicinput *DeleteNetworkInterfaceParams) (result DeleteNetworkInterfaceResult, err error)
	AttachNetworkInterface(ctx context.Context, nicinput *AttachNetworkInterfaceParams) (result AttachNetworkInterfaceResult, err error)
	DetachNetworkInterface(ctx context.Context, nicinput *DetachNetworkInterfaceParams) (result DetachNetworkInterfaceResult, err error)
	CreateFirewall(ctx context.Context, config configuration.FirewallConfiguration) (*configuration.FirewallConfiguration, error)
	
}

// create new client to acess all network functions
func NewNetworkClient() Network {
	return &network{}
}

// network implements Network
type network struct {
	Token string `json:"token"`
	
}

// CreateNetworkInterface creates a new network interface
func (n *network) CreateNetworkInterface(ctx context.Context, nicinput *NetworkInterfaceParams) (result InterfaceResult, err error) {
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
