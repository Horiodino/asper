package configuration

// this will caointain all the configuration such as structs that willpass ti the functons to create appropiate resource

/*
	now we define the erquirements  creating resources  for example , if we create a vm we need configuration
	such as configuration of the vm , we need to know the name of the vm , we need to know the os of the vm
*/

type RunInstanceInput struct {
	Name              asperstring
	OS                asperstring
	SSHConfiguration  InputSSHConfiguration
	NetworkInterface  InputNetworkInterface
	DiskConfiguration InputDiskConfiguration
}

/* the instancesshconfiguration will be used to knew the key that is need to be used to ssh into the vm */
type InputSSHConfiguration struct {
	Keyname string
}

/* the inputnetworkinterface will be used to know the network interface that is need to be used to ssh into the vm */
type InputNetworkInterface struct {
}

/* the inputdiskconfiguration will be used to know the disk configuration that is need to be used to ssh into the vm */
type InputDiskConfiguration struct {
}

type RunInstanceOutput struct {
}

type FirewallConfiguration struct {
	ID    string
	Name  string
	Rules []FirewallRule
}

type FirewallRule struct {
	ID           string
	Name         string
	InboundRule  []InboundRule
	OutboundRule []OutboundRule
}
