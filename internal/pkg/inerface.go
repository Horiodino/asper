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
	List   List
	// Help()
	// VersionVM()
	// ConfigVM() error
	// SSHVM(name string) error
	MetricsVM
	// CheckHealthVM(id string) error
}

func (c *asper) Help() {
	fmt.Println("Running Help")
}

type Creator interface {
	CreateVM(instanceid asperstring) error
	CreateFirewall(instanceid asperstring) error
	CreateNetwork(instanceid asperstring) error
	CreateDisk(instanceid asperstring) *UserContext
	CreateSSH(instanceid asperstring) error
}

func (c *UserContext) CreateVM(instanceid asperstring) *UserContext {
	fmt.Println("Running CreateVM")
	return c
}
func (c *UserContext) CreateFirewall(instanceid asperstring) *UserContext {
	fmt.Println("Running CreateFirewall")
	return c
}
func (c *UserContext) CreateNetwork(instanceid asperstring) *UserContext {
	fmt.Println("Running CreateNetwork")
	return c
}
func (c *UserContext) CreateDisk(instanceid asperstring) *UserContext {
	fmt.Println("Running CreateDisk")
	return c
}
func (c *UserContext) CreateSSH(instanceid asperstring) *UserContext {
	fmt.Println("Running CreateSSH")
	return c
}

type MetricsVM interface {
	GetCPU() error
	GetMemory() error
	GetDisk() error
	GetNetwork() error
}

func (c *UserContext) GetCPU() *UserContext {
	fmt.Println("Running GetCPU")
	return c
}
func (c *UserContext) GetMemory() *UserContext {
	fmt.Println("Running GetMemory")
	return c
}
func (c *UserContext) GetDisk() *UserContext {
	fmt.Println("Running GetDisk")
	return c
}
func (c *UserContext) GetNetwork() *UserContext {
	fmt.Println("Running GetNetwork")
	return c
}

type Modify interface {
	ModifyVM(instanceid asperstring) error
	MoveVM(instanceid asperstring) error
	ResizeVM(instanceid asperstring) error
	RebootVM(instanceid asperstring) error
	StopVM(instanceid asperstring) error
	StartVM(instanceid asperstring) error
	RebuildVM(instanceid asperstring) error
	ReinstallVM(instanceid asperstring) error
	ResetVM(instanceid asperstring) error
	CloneVM(instanceid asperstring) error
	BackupVM(instanceid asperstring) error
	RestoreVM(instanceid asperstring) error
	ConvertVM(instanceid asperstring) error
	ModifyFirewall(instanceid asperstring) error
	ModifyNetwork(instanceid asperstring) error
	ModifyDisk(instanceid asperstring) error
	ModifySSH(instanceid asperstring) error
}

func (c *UserContext) ModifyVM(instanceid asperstring) *UserContext {
	fmt.Println("Running ModifyVM")
	return c
}
func (c *UserContext) MoveVM(instanceid asperstring) *UserContext {
	fmt.Println("Running MoveVM")
	return c
}
func (c *UserContext) ResizeVM(instanceid asperstring) *UserContext {
	fmt.Println("Running ResizeVM")
	return c
}
func (c *UserContext) RebootVM(instanceid asperstring) *UserContext {
	fmt.Println("Running RebootVM")
	return c
}
func (c *UserContext) StopVM(instanceid asperstring) *UserContext {
	fmt.Println("Running StopVM")
	return c
}
func (c *UserContext) StartVM(instanceid asperstring) *UserContext {
	fmt.Println("Running StartVM")
	return c
}
func (c *UserContext) RebuildVM(instanceid asperstring) *UserContext {
	fmt.Println("Running RebuildVM")
	return c
}
func (c *UserContext) ReinstallVM(instanceid asperstring) *UserContext {
	fmt.Println("Running ReinstallVM")
	return c
}
func (c *UserContext) ResetVM(instanceid asperstring) *UserContext {
	fmt.Println("Running ResetVM")
	return c
}
func (c *UserContext) CloneVM(instanceid asperstring) *UserContext {
	fmt.Println("Running CloneVM")
	return c
}
func (c *UserContext) BackupVM(instanceid asperstring) *UserContext {
	fmt.Println("Running BackupVM")
	return c
}
func (c *UserContext) RestoreVM(instanceid asperstring) *UserContext {
	fmt.Println("Running RestoreVM")
	return c
}
func (c *UserContext) ConvertVM(instanceid asperstring) *UserContext {
	fmt.Println("Running ConvertVM")
	return c
}
func (c *UserContext) ModifyFirewall(instanceid asperstring) *UserContext {
	fmt.Println("Running ModifyFirewall")
	return c
}
func (c *UserContext) ModifyNetwork(instanceid asperstring) *UserContext {
	fmt.Println("Running ModifyNetwork")
	return c
}
func (c *UserContext) ModifyDisk(instanceid asperstring) *UserContext {
	fmt.Println("Running ModifyDisk")
	return c
}
func (c *UserContext) ModifySSH(instanceid asperstring) *UserContext {
	fmt.Println("Running ModifySSH")
	return c
}

type Delete interface {
	DeleteVM(instanceid asperstring) error
	DeleteFirewall(instanceid asperstring) error
	DeleteNetwork(instanceid asperstring) error
	DeleteDisk(instanceid asperstring) error
	DeleteSSH(instanceid asperstring) error
}

func (c *UserContext) DeleteVM(instanceid asperstring) *UserContext {
	fmt.Println("Running DeleteVM")
	return c
}
func (c *UserContext) DeleteFirewall(instanceid asperstring) *UserContext {
	fmt.Println("Running DeleteFirewall")
	return c
}
func (c *UserContext) DeleteNetwork(instanceid asperstring) *UserContext {
	fmt.Println("Running DeleteNetwork")
	return c
}
func (c *UserContext) DeleteDisk(instanceid asperstring) *UserContext {
	fmt.Println("Running DeleteDisk")
	return c
}
func (c *UserContext) DeleteSSH(instanceid asperstring) *UserContext {
	fmt.Println("Running DeleteSSH")
	return c
}

type FirewallRules interface {
	AddFirewallRule() error
	DeleteFirewallRule() error
	GetFirewallRule() error
	UpdateFirewallRule() error
	ListFirewallRule() error
}

func (c *UserContext) AddFirewallRule() *UserContext {
	fmt.Println("Running AddFirewallRule")
	return c
}
func (c *UserContext) DeleteFirewallRule() *UserContext {
	fmt.Println("Running DeleteFirewallRule")
	return c
}
func (c *UserContext) GetFirewallRule() *UserContext {
	fmt.Println("Running GetFirewallRule")
	return c
}
func (c *UserContext) UpdateFirewallRule() *UserContext {
	fmt.Println("Running UpdateFirewallRule")
	return c
}
func (c *UserContext) ListFirewallRule() *UserContext {
	fmt.Println("Running ListFirewallRule")
	return c
}

type Get interface {
	InstanceState(instanceid asperstring) error
	InstanceIP(instanceid asperstring) error
	InstanceSSHKey(instanceid asperstring) error
	InstanceNetworkInterface(instanceid asperstring) error
	InstanceDiskConfiguration(instanceid asperstring) error
	InstanceFirewallConfiguration(instanceid asperstring) error
}

func (c *UserContext) GetInstanceState(instanceid asperstring) *UserContext {
	fmt.Println("Running GetInstanceState")
	return c
}
func (c *UserContext) GetInstanceIP(instanceid asperstring) *UserContext {
	fmt.Println("Running GetInstanceIP")
	return c
}
func (c *UserContext) GetInstanceSSHKey(instanceid asperstring) *UserContext {
	fmt.Println("Running GetInstanceSSHKey")
	return c
}
func (c *UserContext) GetInstanceNetworkInterface(instanceid asperstring) *UserContext {
	fmt.Println("Running GetInstanceNetworkInterface")
	return c
}
func (c *UserContext) GetInstanceDiskConfiguration(instanceid asperstring) *UserContext {
	fmt.Println("Running GetInstanceDiskConfiguration")
	return c
}
func (c *UserContext) GetInstanceFirewallConfiguration(instanceid asperstring) *UserContext {
	fmt.Println("Running GetInstanceFirewallConfiguration")
	return c
}

type List interface {
	ListVM() error
	ListFirewall() error
	ListNetwork() error
	ListDisk() error
	ListSSHKey() error
}

func (c *UserContext) ListVM() *UserContext {
	fmt.Println("Running ListVM")
	return c
}
func (c *UserContext) ListFirewall() *UserContext {
	fmt.Println("Running ListFirewall")
	return c
}
func (c *UserContext) ListNetwork() *UserContext {
	fmt.Println("Running ListNetwork")
	return c
}
func (c *UserContext) ListDisk() *UserContext {
	fmt.Println("Running ListDisk")
	return c
}
func (c *UserContext) ListSSHKey() *UserContext {
	fmt.Println("Running ListSSHKey")
	return c
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

	creator := &UserContext{}
	asper := &asper{Create: creator}
	creator.CreateVM("vm1")
}
