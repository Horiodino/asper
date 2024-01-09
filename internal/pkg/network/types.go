package network

import "github.com/horiodino/asper/internal/logger"

type DescribeFirewallConfigurationInput struct {
	InstanceID logger.Asperstring
}

type DescribeFirewallConfigurationOutput struct {
	FirewallConfiguration FirewallConfiguration
}

type NetworkInterfaceParams struct {
	Name   string `json:"name"`
	IP     string `json:"ip"`
	Subnet string `json:"subnet"`
}

type InterfaceResult struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	IP     string `json:"ip"`
	Subnet string `json:"subnet"`
	Status string `json:"status"`
}

type DeleteNetworkInterfaceParams struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type DeleteNetworkInterfaceResult struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"status"`
}

type AttachNetworkInterfaceParams struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type AttachNetworkInterfaceResult struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"status"`
}

type DetachNetworkInterfaceParams struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type DetachNetworkInterfaceResult struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"status"`
}
type BridgeConfigurationInput struct {
}

type BridgeConfigurationOutput struct {
}
