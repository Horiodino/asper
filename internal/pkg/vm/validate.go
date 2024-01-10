package vm

import (
	"fmt"
	"reflect"

	"github.com/horiodino/asper/internal/pkg/ssh"
	"github.com/horiodino/asper/internal/utils"
)

// the validateInput method validates the input for all the methods in the vm package
func validateInput(input interface{}) (string, error) {
	// the validateinput validates based on the input type and returns an error if the input is invalid
	// the types is cheked like the interface is type casted to the specific type and then validated
	// for example we want to have struct for the input of the createVM method
	// so we type cast the input to the struct type and then validate the struct

	// we need to get the values no the type of the input
	// we can get the type of the input using the reflect package

	strtype := reflect.TypeOf(input).String()
	fmt.Println(strtype)

	fmt.Println("Type is ", reflect.TypeOf(input))

	switch strtype {
	case "vm.InstanceConfigurationInput":
		paramerter := &InstanceConfigurationInput{
			Name:              input.(InstanceConfigurationInput).Name,
			OS:                input.(InstanceConfigurationInput).OS,
			Memory:            input.(InstanceConfigurationInput).Memory,
			VCPU:              input.(InstanceConfigurationInput).VCPU,
			Lock:              input.(InstanceConfigurationInput).Lock,
			SSHConfiguration:  input.(InstanceConfigurationInput).SSHConfiguration,
			NetworkInterface:  input.(InstanceConfigurationInput).NetworkInterface,
			DiskConfiguration: input.(InstanceConfigurationInput).DiskConfiguration,
		}
		if paramerter.Name == "" {
			return "", fmt.Errorf("Name field is empty")
		}
		if paramerter.OS == "" {
			return "", fmt.Errorf("OS field is empty")
		}
		if paramerter.Memory == 0 {
			return "", fmt.Errorf("Memory field is empty")
		}
		if paramerter.VCPU == 0 {
			return "", fmt.Errorf("VCPU field is empty")
		}
		if paramerter.SSHConfiguration.Keyname == "" {
			return "", fmt.Errorf("SSHConfiguration.Keyname field is empty")
		}
		if paramerter.NetworkInterface.Name == "" {
			return "", fmt.Errorf("NetworkInterface.Name field is empty")
		}
		if paramerter.NetworkInterface.Bridge == "" {
			return "", fmt.Errorf("NetworkInterface.Bridge field is empty")
		}
		if paramerter.DiskConfiguration.Name == "" {
			return "", fmt.Errorf("DiskConfiguration.Name field is empty")
		}
		if paramerter.DiskConfiguration.Size == 0 {
			return "", fmt.Errorf("DiskConfiguration.Size field is empty")
		}

		err := utils.ValidateName(paramerter.Name, "vm")
		if err != nil {
			return "", err
		}

		resp, err := fetch_system_memory_info(&paramerter.Memory)
		if err != nil {
			return "", err
		}
		if resp == false {
			return "", fmt.Errorf("Insufficient memory")
		}
		resp, err = fetch_system_cpu_info(&paramerter.VCPU)
		if err != nil {
			return "", err
		}
		if resp == false {
			return "", fmt.Errorf("Insufficient CPU")
		}
		resp, err = fetch_system_disk_info(&paramerter.DiskConfiguration.Size)
		if err != nil {
			return "", err
		}

	case "vm.DescribeInstancesInput":
		paramerter := &DescribeInstancesInput{
			Name:   input.(DescribeInstancesInput).Name,
			Filter: input.(DescribeInstancesInput).Filter,
		}
		if paramerter.Name == "" {
			return "", fmt.Errorf("Name field is empty")
		}
		err := utils.ValidateName(paramerter.Name, "vm")
		if err != nil {
			return "", err
		}

		condition := paramerter.Filter.MaxResults
		if condition == 0 && paramerter.Filter.MaxResults == 0 {
			paramerter.Filter.MaxResults = 10
			paramerter.Filter.State = "running"
		}

	case "vm.DescribeSSHKeyInput":
		paramerter := &ssh.DescribeSSHKeyInput{
			Name:       input.(ssh.DescribeSSHKeyInput).Name,
			InstanceID: input.(ssh.DescribeSSHKeyInput).InstanceID,
		}
		if paramerter.Name == "" && paramerter.InstanceID == "" {
			return "", fmt.Errorf("Name and InstanceID field is empty")
		}
	case "vm.DescribeNetworkInterfaceInput":
		paramerter := &DescribeNetworkInterfaceInput{
			InstanceID: input.(DescribeNetworkInterfaceInput).InstanceID,
		}
		if paramerter.InstanceID == "" {
			return "", fmt.Errorf("InstanceID field is empty")
		}
	case "vm.DescribeDiskConfigurationInput":
		paramerter := &DescribeDiskConfigurationInput{
			InstanceID: input.(DescribeDiskConfigurationInput).InstanceID,
		}
		if paramerter.InstanceID == "" {
			return "", fmt.Errorf("InstanceID field is empty")
		}
	case "vm.DescribeFirewallConfigurationInput":
		paramerter := &DescribeFirewallConfigurationInput{
			InstanceID: input.(DescribeFirewallConfigurationInput).InstanceID,
		}
		if paramerter.InstanceID == "" {
			return "", fmt.Errorf("InstanceID field is empty")
		}
	default:
		return "", fmt.Errorf("Invalid input type")
	}
	return "", nil
}

// TODO add logic to fetch the system info
func fetch_system_memory_info(*uint) (bool, error) {
	return false, nil
}
func fetch_system_cpu_info(*uint) (bool, error) {
	return false, nil
}

func fetch_system_disk_info(*uint) (bool, error) {
	return false, nil
}
