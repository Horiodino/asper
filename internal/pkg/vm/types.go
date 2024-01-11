package vm

import (
	"github.com/horiodino/asper/internal/logger"
	"github.com/horiodino/asper/internal/pkg/ssh"
)

type RunInstanceOutput struct {
}

type FirewallConfiguration struct {
	ID    logger.Asperstring `json:"id"`
	Name  logger.Asperstring `json:"name"`
	Rules []FirewallRule     `json:"rules"`
}

type FirewallRule struct {
	ID           logger.Asperstring `json:"id"`
	Name         logger.Asperstring `json:"name"`
	InboundRule  []InboundRule      `json:"inbound_rule"`
	OutboundRule []OutboundRule     `json:"outbound_rule"`
}

type InboundRule struct {
	IPRangeFrom logger.Asperstring `json:"ip_range_from"`
	IPRangeTo   logger.Asperstring `json:"ip_range_to"`
	Protocol    logger.Asperstring `json:"protocol"`
	Ports       []int              `json:"ports"`
}

type OutboundRule struct {
	IPRangeFrom logger.Asperstring `json:"ip_range_from"`
	IPRangeTo   logger.Asperstring `json:"ip_range_to"`
	Protocol    logger.Asperstring `json:"protocol"`
	Ports       []int              `json:"ports"`
}

type InstanceConfigurationInput struct {
	Name              logger.Asperstring        `json:"name"`
	OS                logger.Asperstring        `json:"os"`
	Memory            uint                      `json:"memory"`
	VCPU              uint                      `json:"vcpu"`
	Lock              bool                      `json:"lock"`
	SSHConfiguration  ssh.SSHConfigurationInput `json:"ssh_configuration"`
	NetworkInterface  NetworkInterfaceInput     `json:"network_interface"`
	DiskConfiguration DiskConfigurationInput    `json:"disk_configuration"`
}

type NetworkInterfaceInput struct {
	Name   logger.Asperstring `json:"name"`
	Bridge logger.Asperstring `json:"bridge"`
}
type DiskConfigurationInput struct {
	Name logger.Asperstring `json:"name"`
	Size uint               `json:"size"`
}

type InstanceConfigurationOutput struct {
	ID                    logger.Asperstring          `json:"id"`
	Name                  logger.Asperstring          `json:"name"`
	OS                    logger.Asperstring          `json:"os"`
	SSHConfiguration      ssh.SSHKeyOutput            `json:"ssh_configuration"`
	NetworkInterface      NetworkInterfaceOutput      `json:"network_interface"`
	UserID                logger.Asperstring          `json:"user_id"`
	FirewallConfiguration FirewallConfigurationOutput `json:"firewall_configuration"`
	DiskConfiguration     DiskConfigurationOutput     `json:"disk_configuration"`
}

type NetworkInterfaceOutput struct {
	Name   logger.Asperstring `json:"name"`
	Bridge logger.Asperstring `json:"bridge"`
	Status logger.Asperstring `json:"status"`
	// TODO ADD IP Mac and assigned security group
}

type FirewallConfigurationOutput struct {
	ID         logger.Asperstring   `json:"id"`
	Name       logger.Asperstring   `json:"name"`
	Rules      []FirewallRuleOutput `json:"rules"`
	AppliedFor []logger.Asperstring `json:"applied_for"`
}

type FirewallRuleOutput struct {
	ID           logger.Asperstring   `json:"id"`
	Name         logger.Asperstring   `json:"name"`
	InboundRule  []InboundRuleOutput  `json:"inbound_rule"`
	OutboundRule []OutboundRuleOutput `json:"outbound_rule"`
}

type InboundRuleOutput struct {
	IPRangeFrom logger.Asperstring `json:"ip_range_from"`
	IPRangeTo   logger.Asperstring `json:"ip_range_to"`
	Protocol    logger.Asperstring `json:"protocol"`
	Ports       []int              `json:"ports"`
}

type OutboundRuleOutput struct {
	IPRangeFrom logger.Asperstring `json:"ip_range_from"`
	IPRangeTo   logger.Asperstring `json:"ip_range_to"`
	Protocol    logger.Asperstring `json:"protocol"`
	Ports       []int              `json:"ports"`
}

type DiskConfigurationOutput struct {
	Name logger.Asperstring `json:"name"`
	Size uint               `json:"size"`
}

type DescribeInstancesInput struct {
	Name   logger.Asperstring     `json:"name"`
	Filter DescribeInstanceFilter `json:"filter"`
}

type DescribeInstanceFilter struct {
	State      logger.Asperstring `json:"state"`
	MaxResults uint               `json:"max_results"`
}

type DescribeInstancesOutput struct {
	ResulteMetadata []ResultMetadata `json:"result_metadata"`
}

type ResultMetadata struct {
	ID       logger.Asperstring `json:"id"`
	Name     logger.Asperstring `json:"name"`
	OS       logger.Asperstring `json:"os"`
	SSHKey   logger.Asperstring `json:"ssh_key"`
	Network  logger.Asperstring `json:"network"`
	UserID   logger.Asperstring `json:"user_id"`
	Instance logger.Asperstring `json:"instance"`
}

type DescribeNetworkInterfaceInput struct {
	InstanceID logger.Asperstring `json:"instance_id"`
}

type DescribeNetworkInterfaceOutput struct {
	NetworkInterface NetworkInterfaceOutput `json:"network_interface"`
}

// TODO use slice
type DescribeDiskConfigurationInput struct {
	InstanceID logger.Asperstring `json:"instance_id"`
}

type DescribeDiskConfigurationOutput struct {
	Name       logger.Asperstring `json:"name"`
	Size       uint               `json:"size"`
	Status     logger.Asperstring `json:"status"`
	InstanceID logger.Asperstring `json:"instance_id"`
}

type DescribeFirewallConfigurationInput struct {
	InstanceID logger.Asperstring `json:"instance_id"`
}

type DescribeFirewallConfigurationOutput struct {
	FirewallConfiguration FirewallConfiguration `json:"firewall_configuration"`
}
