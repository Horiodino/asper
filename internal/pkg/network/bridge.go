package network

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/horiodino/asper/bkend/hypervisor/qemu"
)

func (n *network) CreateNetworkBridge(ctx context.Context, nicinput BridgeConfigurationInput) (result BridgeConfigurationOutput, err error) {
	QemuClient := qemu.InitializeQEMU()

	_, err = validateInput(nicinput)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

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
		if parameter.Name == "" {
			return "", fmt.Errorf("Name field is empty")
		}

		if parameter.Bridge == "" {
			return "", fmt.Errorf("Bridge field is empty")
		}

		if parameter.IP == "" {
			return "", fmt.Errorf("IP field is empty")
		}

		if parameter.Mask == "" {
			return "", fmt.Errorf("Mask field is empty")
		}

		// if parameter.IP == "" {
		// 	return "", fmt.Errorf("IP field is empty")
		// }

		// if parameter.Mask == "" {
		// 	return "", fmt.Errorf("Mask field is empty")
		// }

		// err := validateName(parameter.Name, "network")
		// if err != nil {
		// 	return "", err
		// }

		// err = validateIP(parameter.IP)
		// if err != nil {
		// 	return "", err
		// }

		// err = validateMask(parameter.Mask)
		// if err != nil {
		// 	return "", err
		// }

		// return "", nil

	}

	return nil, nil
}
