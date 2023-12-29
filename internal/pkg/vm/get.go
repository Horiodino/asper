package vm

import (
	"fmt"

	"github.com/horiodino/asper/internal/logger"
)

type Get interface {
	InitializeGet() *vm
}

type GetObject interface {
	InstanceState(instanceid logger.Asperstring) *vm
	InstanceIP(instanceid logger.Asperstring) *vm
	InstanceSSHKey(instanceid logger.Asperstring) *vm
	InstanceNetworkInterface(instanceid logger.Asperstring) *vm
	InstanceDiskConfiguration(instanceid logger.Asperstring) *vm
	InstanceFirewallConfiguration(instanceid logger.Asperstring) *vm
}

func (c *vm) InstanceState(instanceid logger.Asperstring) *vm {
	fmt.Println("Running GetInstanceState")
	return c
}
func (c *vm) InstanceIP(instanceid logger.Asperstring) *vm {
	fmt.Println("Running GetInstanceIP")
	return c
}
func (c *vm) InstanceSSHKey(instanceid logger.Asperstring) *vm {
	fmt.Println("Running GetInstanceSSHKey")
	return c
}
func (c *vm) InstanceNetworkInterface(instanceid logger.Asperstring) *vm {
	fmt.Println("Running GetInstanceNetworkInterface")
	return c
}
func (c *vm) InstanceDiskConfiguration(instanceid logger.Asperstring) *vm {
	fmt.Println("Running GetInstanceDiskConfiguration")
	return c
}
func (c *vm) InstanceFirewallConfiguration(instanceid logger.Asperstring) *vm {
	fmt.Println("Running GetInstanceFirewallConfiguration")
	return c
}
