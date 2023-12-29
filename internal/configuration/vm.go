package configuration

import (
	// str "github.com/horiodino/asper/internal/pkg"
	"github.com/horiodino/asper/internal/logger"
)

// this will caointain all the configuration such as structs that willpass ti the functons to create appropiate resource

/*
	now we define the erquirements  creating resources  for example , if we create a vm we need configuration
	such as configuration of the vm , we need to know the name of the vm , we need to know the os of the vm
*/

type RunInstanceInput struct {
	Name              logger.Asperstring
	OS                logger.Asperstring
	SSHConfiguration  InputSSHConfiguration
	NetworkInterface  InputNetworkInterface
	DiskConfiguration InputDiskConfiguration
}

/* the instancesshconfiguration will be used to knew the key that is need to be used to ssh into the vm */
type InputSSHConfiguration struct {
	Keyname logger.Asperstring
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
