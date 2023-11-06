package pkg

import (
	"fmt"
)

type UserContext struct {
	Username string
	Password string
}

type asper struct {
	Create Creator
	Delete Delete
	Get    Get
	Modify Modify
	List   ListOfResources
	// Help()
	// VersionVM()
	// ConfigVM() *UserContext
	// SSHVM(name string) *UserContext
	MetricsVM MetricsVM
	// CheckHealthVM(id string) *UserContext
}

func (c *asper) Help() {
	fmt.Println("Running Help")
}

type Cli struct {
	isadmin bool
}

func NewLocalClient() *Cli {
	fmt.Println("Running NewLocalClient")
	return &Cli{isadmin: true}
}

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

type asperstring string

func String(s string) asperstring {
	return asperstring(s)
}

type asperint int

func Int(i int) asperint {
	return asperint(i)
}

type InboundRule struct {
	iprangeFrom asperstring
	iprangeTo   asperstring
	protocol    asperstring
	ports       []asperint
}

type OutboundRule struct {
}

func main() {

	metrice := &metrice{UserID: "1", IsUSer: true, InstanceID: "1"}
	asper := &asper{MetricsVM: metrice}
	val := asper.Modify.InitilizeModifyClient("")
	val.ModifyVM("").UpdateFirewallRule()

}
