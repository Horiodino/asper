package pkg

import "fmt"

type create struct {
	UserID asperstring
	IsUSer bool
}

type Creator interface {
	InitializeCreateClient() *create
}

func (c *create) InitializeCreateClient() *create {
	fmt.Println("Running InitializeCreateClient")
	return c
}

type Create interface {
	CreateVM(instanceid asperstring) *create
	CreateFirewall(instanceid asperstring) *create
	InitializeNetqorkClient() *network
	CreateDisk(instanceid asperstring) *create
	CreateSSH(instanceid asperstring) *create
}

func (c *create) CreateVM(instanceid asperstring) *create {
	fmt.Println("Running CreateVM")
	return c
}
func (c *create) CreateFirewall(instanceid asperstring) *create {
	fmt.Println("Running CreateFirewall")
	return c
}
func (c *network) InitializeNetqorkClient() *network {
	fmt.Println("Running CreateNetwork")
	return c
}
func (c *create) CreateDisk(instanceid asperstring) *create {
	fmt.Println("Running CreateDisk")
	return c
}
func (c *create) CreateSSH(instanceid asperstring) *create {
	fmt.Println("Running CreateSSH")
	return c
}
