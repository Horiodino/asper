package vm

import (
	"github.com/horiodino/asper/internal/logger"
)

type RunInstanceInput struct {
	Name              logger.Asperstring
	OS                logger.Asperstring
	SSHConfiguration  InputSSHConfiguration
	NetworkInterface  InputNetworkInterface
	DiskConfiguration InputDiskConfiguration
}

type InputSSHConfiguration struct {
	Keyname logger.Asperstring
}

type InputNetworkInterface struct {
}

type InputDiskConfiguration struct {
}

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
	SSHConfiguration  InputSSHConfiguration
	NetworkInterface  InputNetworkInterface
	DiskConfiguration InputDiskConfiguration
}

type InstanceConfigurationOutput struct {
	ID                    logger.Asperstring
	Name                  logger.Asperstring
	OS                    logger.Asperstring
	SSHConfiguration      InputSSHConfiguration
	NetworkInterface      InputNetworkInterface
	UserID                logger.Asperstring
	FirewallConfiguration FirewallConfiguration
	DiskConfiguration     InputDiskConfiguration
}

type DescribeInstancesInput struct {
	InstanceID logger.Asperstring
}

type DescribeInstancesOutput struct {
	InstanceConfigurationOutput InstanceConfigurationOutput
}

type DescribeSSHKeyInput struct {
	InstanceID logger.Asperstring
}

type DescribeSSHKeyOutput struct {
	SSHConfiguration InputSSHConfiguration
}

type DescribeNetworkInterfaceInput struct {
	InstanceID logger.Asperstring
}

type DescribeNetworkInterfaceOutput struct {
	NetworkInterface InputNetworkInterface
}

type DescribeDiskConfigurationInput struct {
	InstanceID logger.Asperstring
}

type DescribeDiskConfigurationOutput struct {
	DiskConfiguration InputDiskConfiguration
}

type DescribeFirewallConfigurationInput struct {
	InstanceID logger.Asperstring
}

type DescribeFirewallConfigurationOutput struct {
	FirewallConfiguration FirewallConfiguration
}

type ALL struct {
	InstanceConfigurationInput  InstanceConfigurationInput
	InstanceConfigurationOutput InstanceConfigurationOutput
	InputDiskConfiguration      InputDiskConfiguration
	InputNetworkInterface       InputNetworkInterface
	InputSSHConfiguration       InputSSHConfiguration
	RunInstanceInput            RunInstanceInput
	RunInstanceOutput           RunInstanceOutput
	FirewallConfiguration       FirewallConfiguration
	FirewallRule                FirewallRule
	InboundRule                 InboundRule
	OutboundRule                OutboundRule
}
