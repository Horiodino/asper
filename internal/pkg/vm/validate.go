package vm

import (
	"fmt"
	"reflect"
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
		isempty, err := paramerter.IsEmpty()
		if err != nil {
			return "", err
		}
		if !isempty {
			return "", fmt.Errorf("missing required parameter")
		}
		return "valid", nil

	case "vm.DescribeInstancesInput":
		paramerter := &DescribeInstancesInput{
			Name:   input.(DescribeInstancesInput).Name,
			Filter: input.(DescribeInstancesInput).Filter,
		}
		isempty, err := paramerter.IsEmpty()
		if err != nil {
			return "", err
		}
		if !isempty {
			return "", fmt.Errorf("missing required parameter")
		}

		return "valid", nil
	}
	return "valid", nil
}

func (i *InstanceConfigurationInput) IsEmpty() (bool, error) {
	paramerter := []string{i.Name, i.OS, string(rune(i.Memory)), string(rune(i.VCPU))}
	for _, value := range paramerter {
		if value == "" {
			return false, fmt.Errorf("missing required parameter: %s", reflect.TypeOf(value))
		}
	}

	if i.SSHConfiguration.Keyname == "" {
		return false, fmt.Errorf("missing required parameter: %s", i.SSHConfiguration.Keyname)
	} else {
		// Set it true in DB
		i.SSHConfiguration.Keyname = "true"
	}

	paramerter = []string{i.NetworkInterface.IP, i.NetworkInterface.Name, i.NetworkInterface.Subnet}
	for _, value := range paramerter {
		if value == "" {
			return false, fmt.Errorf("missing required parameter: %s", value)
		}
	}

	paramerter = []string{i.DiskConfiguration.Name, string(rune(i.DiskConfiguration.Size))}
	for _, value := range paramerter {
		if value == "" {
			return false, fmt.Errorf("missing required parameter: %s", value)
		}
	}

	return true, nil
}

func (i *DescribeInstancesInput) IsEmpty() (bool, error) {
	paramerter := []string{i.Name, i.Filter.State}
	for _, value := range paramerter {
		if value == "" {
			return false, fmt.Errorf("missing required parameter: %s", value)
		}
	}

	if i.Filter.State == "" {
		return false, fmt.Errorf("missing required parameter: %s", i.Filter.State)
	}
	if i.Filter.MaxResults == 0 {
		return false, fmt.Errorf("missing required parameter: %d", i.Filter.MaxResults)
	}
	return true, nil
}

//  a global function to tell if any of the input is empty
// it will take any interface as input and return error if any of the input is empty

// first we marshal the input to json and in json , we see coressponding to the empty input
// if it empty like OS : "" then we do something see the nil
// ok what is nill? we seee coresponing value , ok its OS
// sos we print the error that OS is empty

// func isEmpty(input InstanceConfigurationInput) error {
// 	// marshal the input to json
// 	jsondata, err := json.Marshal(input)
// 	if err != nil {
// 		return err
// 	}

// 	// unmarshal the json to map
// 	var data map[string]interface{}
// 	err = json.Unmarshal(jsondata, &data)
// 	if value == nil {
// 		return fmt.Errorf("missing required parameter: %s", key)
// 	}

// 	return nil
// }
