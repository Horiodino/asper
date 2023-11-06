package pkg

import "fmt"

func (f *network) CreateSubnet() *network {
	fmt.Println("Running CreateSubnet")
	return f
}
func (f *network) CreateRouter() *network {
	fmt.Println("Running CreateRouter")
	return f
}

type network struct {
	UserID asperstring
	IsUSer bool
}

type disk struct {
	UserID asperstring
	IsUSer bool
}

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
