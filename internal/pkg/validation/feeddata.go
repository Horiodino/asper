package validation_test

import "fmt"

type Operation interface {
	Create(service string) (bool, error)
	List() (bool, error)
	Delete() (bool, error)
	Modify() (bool, error)
}

func (c *CliValidator) Create(service string) (bool, error) {
	switch service {
	case "vm":
		fmt.Println("Running CreateVM-validation")
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
