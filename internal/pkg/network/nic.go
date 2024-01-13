package network

import (
	"context"
)

// Network the newtwork package allocates networking resources required
type network struct{}

type Network interface {
	CreateNetworkInterface(ctx context.Context, nicinput NetworkInterfaceParams) (result *InterfaceResult, err error)
	DeleteNetworkInterface(ctx context.Context, nicinput DeleteNetworkInterfaceParams) (result *DeleteNetworkInterfaceParams, err error)
	AttachNetworkInterface(ctx context.Context, nicinput AttachNetworkInterfaceParams) (result *AttachNetworkInterfaceResult, err error)
	DetachNetworkInterface(ctx context.Context, nicinput DetachNetworkInterfaceParams) (result *DetachNetworkInterfaceResult, err error)
	CreateFirewall(ctx context.Context, nicinput NetworkInterfaceParams) (FirewallConfigurationOutput, err error)
	DeleteFirewall(ctx context.Context, nicinput NetworkInterfaceParams) (FirewallConfigurationOutput, err error)
	CheckStatus(ctx context.Context, nicinput NetworkInterfaceParams) (result *InterfaceResult, err error)
	CreateNetworkBridge(ctx context.Context, input BridgeConfigurationInput) (result *BridgeConfigurationOutput, err error)
	DeleteNetworkBridge(ctx context.Context, input BridgeConfigurationInput) (result *BridgeConfigurationOutput, err error)
}

func NewNetworkClient() *network {
	return &network{}
}

func (n *network) CreateNetworkInterface(ctx context.Context, nicinput NetworkInterfaceParams) (result *InterfaceResult, err error) {
	return nil, nil
}

func (n *network) DeleteNetworkInterface(ctx context.Context, nicinput DeleteNetworkInterfaceParams) (result *DeleteNetworkInterfaceParams, err error) {
	return nil, nil
}

func (n *network) AttachNetworkInterface(ctx context.Context, nicinput AttachNetworkInterfaceParams) (result *AttachNetworkInterfaceResult, err error) {
	return nil, nil
}

func (n *network) DetachNetworkInterface(ctx context.Context, nicinput DetachNetworkInterfaceParams) (result *DetachNetworkInterfaceResult, err error) {
	return nil, nil
}

func (n *network) CheckStatus(ctx context.Context, nicinput NetworkInterfaceParams) (result *InterfaceResult, err error) {
	return nil, nil
}
