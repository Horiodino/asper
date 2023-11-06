package pkg

import "fmt"

type Delete interface {
	InitilizeDeleteClient() *delete
}

func (d *delete) InitilizeDeleteClient(instanceid asperstring) *delete {
	fmt.Println("Running DeleteVM")
	return d
}

type deleteresource interface {
	DeleteVM(instanceid asperstring) *delete
	DeleteFirewall(instanceid asperstring) *delete
	DeleteNetwork(instanceid asperstring) *delete
	DeleteDisk(instanceid asperstring) *delete
	DeleteSSH(instanceid asperstring) *delete
}

type delete struct {
	UserID asperstring
	IsUSer bool
}

func (c *delete) DeleteVM(instanceid asperstring) *delete {
	fmt.Println("Running DeleteVM")
	return c
}
func (c *delete) DeleteFirewall(instanceid asperstring) *delete {
	fmt.Println("Running DeleteFirewall")
	return c
}
func (c *delete) DeleteNetwork(instanceid asperstring) *delete {
	fmt.Println("Running DeleteNetwork")
	return c
}
func (c *delete) DeleteDisk(instanceid asperstring) *delete {
	fmt.Println("Running DeleteDisk")
	return c
}
func (c *delete) DeleteSSH(instanceid asperstring) *delete {
	fmt.Println("Running DeleteSSH")
	return c
}
