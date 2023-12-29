package network

type NetworkCreator interface {
	CreateSubnet() *network
	CreateRouter() *network
	CreateLoadBalancer() *network
	CreateFloatingIP() *network
	CreateSecurityGroup() *network
	CreateSecurityGroupRule() *network
	CreatePort() *network
	CreateDNS() *network
	CreateVPN() *network
	CreateIPSec() *network
	CreateIPSecPolicy() *network
	CreateIPSecSiteConnection() *network
	CreateIPSecIKEPolicy() *network
	CreateIPSecEndpointGroup() *network
	CreateIPSecPeerCidr() *network
	CreateIPSecPSK() *network
	CreateIPSecCertificate() *network
	CreateIPSecCertificateRequest() *network
	CreateIPSecCertificateAuthority() *network
	CreateIPSecCertificateAuthorityRequest() *network
	CreateIPSecCertificateAuthorityCRL() *network
	CreateIPSecCertificateAuthorityCRLRequest() *network
	CreateIPSecCertificateAuthorityCRLDistributionPoint() *network
	CreateIPSecCertificateAuthorityOCSP() *network
	CreateIPSecCertificateAuthorityOCSPRequest() *network
	CreateIPSecCertificateAuthorityOCSPResponder() *network
	CreateIPSecCertificateAuthorityOCSPResponderRequest() *network
	CreateIPSecCertificateAuthorityOCSPResponderCertificate() *network
	CreateIPSecCertificateAuthorityOCSPResponderCertificateRequest() *network
	CreateIPSecCertificateAuthorityOCSPResponderCertificateAuthority() *network
	CreateIPSecCertificateAuthorityOCSPResponderCertificateAuthorityRequest() *network
	CreateIPSecCertificateAuthorityOCSPResponderCertificateAuthorityCRL() *network
	CreateIPSecCertificateAuthorityOCSPResponderCertificateAuthorityCRLRequest() *network
	CreateIPSecCertificateAuthorityOCSPResponderCertificateAuthorityCRLDistributionPoint() *network
	CreateIPSecCertificateAuthorityOCSPResponderCertificateAuthorityOCSP() *network
	CreateIPSecCertificateAuthorityOCSPResponderCertificateAuthorityOCSPRequest() *network
}

type NetworkInterfaceParams struct {
	Name   string `json:"name"`   // Name of the network interface to create
	IP     string `json:"ip"`     // IP address to assign to the new network interface
	Subnet string `json:"subnet"` // Subnet to assign to the new network interface
}

type InterfaceResult struct {
	Name   string `json:"name"`   // Name of the network interface
	ID     string `json:"id"`     // ID of the interface
	IP     string `json:"ip"`     // IP address of the network interface
	Subnet string `json:"subnet"` // Subnet of the network interface
	Status string `json:"status"` // Status of the network interface
}

type DeleteNetworkInterfaceParams struct {
	Name string `json:"name"` // Name of the network interface to delete
	ID   string `json:"id"`   // ID of the network interface to delete
}

type DeleteNetworkInterfaceResult struct {
	Name   string `json:"name"`   // Name of the network interface
	ID     string `json:"id"`     // ID of the interface
	Status string `json:"status"` // Status of the network interface
}

type AttachNetworkInterfaceParams struct {
	Name string `json:"name"` // Name of the network interface to attach
	ID   string `json:"id"`   // ID of the network interface to attach
}

type AttachNetworkInterfaceResult struct {
	Name   string `json:"name"`   // Name of the network interface
	ID     string `json:"id"`     // ID of the interface
	Status string `json:"status"` // Status of the network interface
}

type DetachNetworkInterfaceParams struct {
	Name string `json:"name"` // Name of the network interface to detach
	ID   string `json:"id"`   // ID of the network interface to detach
}

type DetachNetworkInterfaceResult struct {
	Name   string `json:"name"`   // Name of the network interface
	ID     string `json:"id"`     // ID of the interface
	Status string `json:"status"` // Status of the network interface
}
