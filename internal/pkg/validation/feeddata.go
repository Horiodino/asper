package validation_test

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/horiodino/asper/internal/configuration"
	"github.com/horiodino/asper/internal/logger"
	"github.com/horiodino/asper/internal/pkg"
	"github.com/horiodino/asper/validator"
)

type Operation interface {
	Create(service string) (bool, error)
	List() (bool, error)
	Delete() (bool, error)
	Modify() (bool, error)
	Client() (bool, error)
}

// this struct will coinatin the client that will be used to validate the input and the all data to be used
// it will be parsed as pointer to the original struct
type Request struct {
	Token         string             `json:"token"`
	RequestType   string             `json:"request_type"`
	Service       string             `json:"service"`
	Configuration *configuration.ALL `json:"configuration"`
}

var User_Request = &Request{}

func (c *CliValidator) Create(service string) (bool, error) {

	Client := pkg.AsperClient()
	v_client := validator.InitializeValidator(c.request.Configuration)

	switch c.request.Service {
	case "vm":
		create_client := Client.Create.InitializeCreateClient()
		resource, err := v_client.IsValidfields(context.Background(), validator.InstanceConfigurationInput)
		if err != nil {
			fmt.Println(err)
		}
		// add new constraints name managed or unmanaged
		// if managed all thing will be allocated by the system and the user will not have to worry about anything
		// allocate resources
		// modify the request
		create_client.CreateVM(context.Background(), resource, "managed")

		// print instance inputconfiguration
		jsondata := c.request.Configuration.InstanceConfigurationInput
		jsondatajson, err := json.Marshal(jsondata)
		if err != nil {
			fmt.Println(err)
		}
		logger := logger.NewLogger()
		logger.LogInfo(string(jsondatajson))

	case "firewall":
		fmt.Println("Running CreateFirewall")
	case "network":
		fmt.Println("Running CreateNetwork")
	case "disk":
		fmt.Println("Running CreateDisk")
	case "ssh":
		fmt.Println("Running CreateSSH")
	default:
		fmt.Println("Invalid service")
	}
	return true, nil
}

func (c *CliValidator) List() (bool, error) {
	fmt.Println("Running List")
	return true, nil
}

func (c *CliValidator) Delete() (bool, error) {
	fmt.Println("Running Delete")
	return true, nil
}

func (c *CliValidator) Modify() (bool, error) {
	fmt.Println("Running Modify")
	return true, nil
}
