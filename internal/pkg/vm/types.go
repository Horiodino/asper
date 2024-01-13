package vm

import (
	"github.com/horiodino/asper/internal/pkg/disk"
	"github.com/horiodino/asper/internal/pkg/network"
	"github.com/horiodino/asper/internal/pkg/ssh"
)

type InstanceConfigurationInput struct {
	Name              string                         `json:"name"`
	OS                string                         `json:"os"`
	Memory            uint                           `json:"memory"`
	VCPU              uint                           `json:"vcpu"`
	Lock              bool                           `json:"lock"`
	SSHConfiguration  ssh.SSHConfigurationInput      `json:"ssh_configuration"`
	NetworkInterface  network.NetworkInterfaceParams `json:"network_interface"`
	DiskConfiguration disk.DiskConfigurationInput    `json:"disk_configuration_input"`
}

type InstanceConfigurationOutput struct {
	ID                    string                              `json:"id"`
	Name                  string                              `json:"name"`
	OS                    string                              `json:"os"`
	SSHConfiguration      ssh.SSHKeyOutput                    `json:"ssh_configuration"`
	NetworkInterface      network.InterfaceResult             `json:"network_interface_result"`
	UserID                string                              `json:"user_id"`
	FirewallConfiguration network.FirewallConfigurationOutput `json:"firewall_configuration"`
	DiskConfiguration     disk.DiskConfigurationOutput        `json:"disk_configuration_output"`
}

type DescribeInstancesInput struct {
	Name   string                 `json:"name"`
	Filter DescribeInstanceFilter `json:"filter"`
}

type DescribeInstanceFilter struct {
	State      string `json:"state"`
	MaxResults uint   `json:"max_results"`
}

type DescribeInstancesOutput struct {
	ResulteMetadata []ResultMetadata `json:"result_metadata"`
}

type ResultMetadata struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	OS       string `json:"os"`
	SSHKey   string `json:"ssh_key"`
	Network  string `json:"network"`
	UserID   string `json:"user_id"`
	Instance string `json:"instance"`
}
