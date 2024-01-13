package network

import (
	"fmt"
	"reflect"
)

// the validateInput method validates the input for all the methods in the network package
func validateInput(input interface{}) (string, error) {
	objType := reflect.TypeOf(input).String()
	fmt.Println(objType)
	fmt.Println("Type is ", reflect.TypeOf(input))
	switch objType {
	case "network.BridgeConfigurationInput":
		parameter := &BridgeConfigurationInput{
			Name:   input.(BridgeConfigurationInput).Name,
			Bridge: input.(BridgeConfigurationInput).Bridge,
			IP:     input.(BridgeConfigurationInput).IP,
			Mask:   input.(BridgeConfigurationInput).Mask,
		}
		isempty, err := parameter.IsEmpty()
		if err != nil {
			return "", err
		}
		if !isempty {
			return "", fmt.Errorf("missing required parameter")
		}
		return "bridge", nil
	case "network.DetachNetworkInterfaceParams":
		parameter := &DetachNetworkInterfaceParams{
			Name: input.(DetachNetworkInterfaceParams).Name,
			ID:   input.(DetachNetworkInterfaceParams).ID,
		}
		err := parameter.IsEmpty()
		if err != nil {
			return "", err
		}
		return "detach", nil
	case "network.AttachNetworkInterfaceParams":
		parameter := &AttachNetworkInterfaceParams{
			Name: input.(AttachNetworkInterfaceParams).Name,
			ID:   input.(AttachNetworkInterfaceParams).ID,
		}
		err := parameter.IsEmpty()
		if err != nil {
			return "", err
		}
		return "attach", nil
	case "network.DeleteNetworkInterfaceParams":
		parameter := &DeleteNetworkInterfaceParams{
			Name: input.(DeleteNetworkInterfaceParams).Name,
			ID:   input.(DeleteNetworkInterfaceParams).ID,
		}
		err := parameter.IsEmpty()
		if err != nil {
			return "", err
		}
		return "delete", nil
	case "network.NetworkInterfaceParams":
		parameter := &NetworkInterfaceParams{
			Name:   input.(NetworkInterfaceParams).Name,
			IP:     input.(NetworkInterfaceParams).IP,
			Subnet: input.(NetworkInterfaceParams).Subnet,
		}
		err := parameter.IsEmpty()
		if err != nil {
			return "", err
		}
		return "create", nil
	case "network.DescribeFirewallConfigurationInput":
		parameter := &DescribeFirewallConfigurationInput{
			InstanceID: input.(DescribeFirewallConfigurationInput).InstanceID,
		}
		err := parameter.IsEmpty()
		if err != nil {
			return "", err
		}
		return "describe", nil

	default:
		return "", fmt.Errorf("invalid input type")
	}
}

func (b *BridgeConfigurationInput) IsEmpty() (bool, error) {
	parameters := []string{b.Name, b.Bridge, b.IP, b.Mask}
	for _, parameter := range parameters {
		if parameter == "" {
			return true, fmt.Errorf("missing required parameter")
		}
	}
	return false, nil
}

func (d *DetachNetworkInterfaceParams) IsEmpty() error {
	parameters := []string{d.Name, d.ID}
	for _, parameter := range parameters {
		if parameter == "" {
			return fmt.Errorf("missing required parameter")
		}
	}
	return nil
}

func (a *AttachNetworkInterfaceParams) IsEmpty() error {
	parameters := []string{a.Name, a.ID}
	for _, parameter := range parameters {
		if parameter == "" {
			return fmt.Errorf("missing required parameter")
		}
	}
	return nil
}

func (d *DeleteNetworkInterfaceParams) IsEmpty() error {
	parameters := []string{d.Name, d.ID}
	for _, parameter := range parameters {
		if parameter == "" {
			return fmt.Errorf("missing required parameter")
		}
	}
	return nil
}

func (n *NetworkInterfaceParams) IsEmpty() error {
	parameters := []string{n.Name, n.IP, n.Subnet}
	for _, parameter := range parameters {
		if parameter == "" {
			return fmt.Errorf("missing required parameter")
		}
	}
	return nil
}

func (d *DescribeFirewallConfigurationInput) IsEmpty() error {
	if d.InstanceID == "" {
		return fmt.Errorf("missing required parameter")
	}
	return nil
}
