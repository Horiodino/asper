package vm

import (
	"github.com/horiodino/asper/internal/logger"
)

type RunInstanceOutput struct {
}

type FirewallConfiguration struct {
	ID    logger.Asperstring
	Name  logger.Asperstring
	Rules []FirewallRule
}

type FirewallRule struct {
	ID           logger.Asperstring
	Name         logger.Asperstring
	InboundRule  []InboundRule
	OutboundRule []OutboundRule
}

type InboundRule struct {
	IPRangeFrom logger.Asperstring
	IPRangeTo   logger.Asperstring
	Protocol    logger.Asperstring
	Ports       []int
}

type OutboundRule struct {
	IPRangeFrom logger.Asperstring
	IPRangeTo   logger.Asperstring
	Protocol    logger.Asperstring
	Ports       []int
}

type InstanceConfigurationInput struct {
	Name              logger.Asperstring
	OS                logger.Asperstring
	Memory            uint
	VCPU              uint
	Lock              bool
	SSHConfiguration  SSHConfigurationInput
	NetworkInterface  NetworkInterfaceInput
	DiskConfiguration DiskConfigurationInput
}
type SSHConfigurationInput struct {
	Keyname logger.Asperstring
}
type NetworkInterfaceInput struct {
	Name   logger.Asperstring
	Bridge logger.Asperstring
}
type DiskConfigurationInput struct {
	Name logger.Asperstring
	Size uint
}

type InstanceConfigurationOutput struct {
	ID                    logger.Asperstring
	Name                  logger.Asperstring
	OS                    logger.Asperstring
	SSHConfiguration      SSHKeyOutput
	NetworkInterface      NetworkInterfaceOutput
	UserID                logger.Asperstring
	FirewallConfiguration FirewallConfigurationOutput
	DiskConfiguration     DiskConfigurationOutput
}

type SSHKeyOutput struct {
	Keyname logger.Asperstring
}
type NetworkInterfaceOutput struct {
	Name   logger.Asperstring
	Bridge logger.Asperstring
	Status logger.Asperstring
	// TODO ADD IP Mac and assigned security group
}

type FirewallConfigurationOutput struct {
	ID         logger.Asperstring
	Name       logger.Asperstring
	Rules      []FirewallRuleOutput
	AppliedFor []logger.Asperstring
}

type FirewallRuleOutput struct {
	ID           logger.Asperstring
	Name         logger.Asperstring
	InboundRule  []InboundRuleOutput
	OutboundRule []OutboundRuleOutput
}

type InboundRuleOutput struct {
	IPRangeFrom logger.Asperstring
	IPRangeTo   logger.Asperstring
	Protocol    logger.Asperstring
	Ports       []int
}

type OutboundRuleOutput struct {
	IPRangeFrom logger.Asperstring
	IPRangeTo   logger.Asperstring
	Protocol    logger.Asperstring
	Ports       []int
}

type DiskConfigurationOutput struct {
	Name logger.Asperstring
	Size uint
}

type DescribeInstancesInput struct {
	Name   logger.Asperstring
	Filter []DescribeInstanceFilter
}

type DescribeInstanceFilter struct {
	State logger.Asperstring
}

type DescribeInstancesOutput struct {
	ResulteMetadata []ResultMetadata
}

type ResultMetadata struct {
	ID       logger.Asperstring
	Name     logger.Asperstring
	OS       logger.Asperstring
	SSHKey   logger.Asperstring
	Network  logger.Asperstring
	UserID   logger.Asperstring
	Instance logger.Asperstring
}

type DescribeSSHKeyInput struct {
	Name logger.Asperstring
	// only any one field is valid
	InstanceID logger.Asperstring
}

type DescribeSSHKeyOutput struct {
	Keyname logger.Asperstring
	UsedBy  []logger.Asperstring
}

type DescribeNetworkInterfaceInput struct {
	InstanceID logger.Asperstring
}

type DescribeNetworkInterfaceOutput struct {
	NetworkInterface NetworkInterfaceOutput
}

// TODO use slice
type DescribeDiskConfigurationInput struct {
	InstanceID logger.Asperstring
}

type DescribeDiskConfigurationOutput struct {
	Name       logger.Asperstring
	Size       uint
	Status     logger.Asperstring
	InstanceID logger.Asperstring
}

type DescribeFirewallConfigurationInput struct {
	InstanceID logger.Asperstring
}

type DescribeFirewallConfigurationOutput struct {
	FirewallConfiguration FirewallConfiguration
}
