package vm

import (
	"context"
)

func (c *vm) DescribeInstances(ctx context.Context, input *DescribeInstancesInput) (*DescribeInstancesOutput, error) {
	return &DescribeInstancesOutput{}, nil
}
func (c *vm) DescribeSSHKey(ctx context.Context, input *DescribeSSHKeyInput) (*DescribeSSHKeyOutput, error) {
	return &DescribeSSHKeyOutput{}, nil
}
func (c *vm) DescribeNetworkInterface(ctx context.Context, input *DescribeNetworkInterfaceInput) (*DescribeNetworkInterfaceOutput, error) {
	return &DescribeNetworkInterfaceOutput{}, nil
}
func (c *vm) DescribeDiskConfiguration(ctx context.Context, input *DescribeDiskConfigurationInput) (*DescribeDiskConfigurationOutput, error) {
	return &DescribeDiskConfigurationOutput{}, nil
}
func (c *vm) DescribeFirewallConfiguration(ctx context.Context, input *DescribeFirewallConfigurationInput) (*DescribeFirewallConfigurationOutput, error) {
	return &DescribeFirewallConfigurationOutput{}, nil
}
