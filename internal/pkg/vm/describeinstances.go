package vm

import (
	"context"
	"fmt"
)

func (c *vm) DescribeInstances(ctx context.Context, input *DescribeInstancesInput) *DescribeInstancesOutput {
	fmt.Println("Running GetInstanceState")
	return nil
}
func (c *vm) DescribeSSHKey(ctx context.Context, input DescribeSSHKeyInput) *DescribeSSHKeyOutput {
	fmt.Println("Running GetInstanceSSHKey")
	return nil
}
func (c *vm) DescribeNetworkInterface(ctx context.Context, input *DescribeNetworkInterfaceInput) *DescribeNetworkInterfaceOutput {
	fmt.Println("Running GetInstanceNetworkInterface")
	return nil
}
func (c *vm) DescribeDiskConfiguration(ctx context.Context, input *DescribeDiskConfigurationInput) *DescribeDiskConfigurationOutput {
	fmt.Println("Running GetInstanceDiskConfiguration")
	return nil
}
func (c *vm) DescribeFirewallConfiguration(ctx context.Context, input *DescribeFirewallConfigurationInput) *DescribeFirewallConfigurationOutput {
	fmt.Println("Running GetInstanceFirewallConfiguration")
	return nil
}
