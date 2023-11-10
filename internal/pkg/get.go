package pkg

import "fmt"

type get struct {
	UserID   asperstring
	Password asperstring
}

type Get interface {
	InitializeGet() *get
}

type GetObject interface {
	InstanceState(instanceid asperstring) *get
	InstanceIP(instanceid asperstring) *get
	InstanceSSHKey(instanceid asperstring) *get
	InstanceNetworkInterface(instanceid asperstring) *get
	InstanceDiskConfiguration(instanceid asperstring) *get
	InstanceFirewallConfiguration(instanceid asperstring) *get
}

func (g *get) InitializeGet() *get {
	fmt.Println("Running InitializeGet")
	return g
}

func (c *get) InstanceState(instanceid asperstring) *get {
	fmt.Println("Running GetInstanceState")
	return c
}
func (c *get) InstanceIP(instanceid asperstring) *get {
	fmt.Println("Running GetInstanceIP")
	return c
}
func (c *get) InstanceSSHKey(instanceid asperstring) *get {
	fmt.Println("Running GetInstanceSSHKey")
	return c
}
func (c *get) InstanceNetworkInterface(instanceid asperstring) *get {
	fmt.Println("Running GetInstanceNetworkInterface")
	return c
}
func (c *get) InstanceDiskConfiguration(instanceid asperstring) *get {
	fmt.Println("Running GetInstanceDiskConfiguration")
	return c
}
func (c *get) InstanceFirewallConfiguration(instanceid asperstring) *get {
	fmt.Println("Running GetInstanceFirewallConfiguration")
	return c
}
