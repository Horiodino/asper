package network

type DescribeFirewallConfigurationInput struct {
	InstanceID string
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
	Name   string `json:"name"`
	Bridge string `json:"bridge"`
	IP     string `json:"ip"`
	Mask   string `json:"mask"`
}

type BridgeConfigurationOutput struct {
}

type FirewallConfigurationInput struct {
	InstanceID         string                        `json:"instance_id"`
	ConfigurationInput *[]FirewallConfigurationInput `json:"configuration_input"`
}

type FirewallConfiguration struct {
	Protocol string `json:"protocol"`
	Port     string `json:"port"`
	SourceIP string `json:"source_ip"`
	DestIP   string `json:"dest_ip"`
}

type FirewallConfigurationOutput struct {
	InstanceID string                   `json:"instance_id"`
	Rules      *[]FirewallConfiguration `json:"rules"`
}
