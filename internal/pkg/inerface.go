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
	res := val.ModifyVM("")
	res.AddFirewallRule()

}
