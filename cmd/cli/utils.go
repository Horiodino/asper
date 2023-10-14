package cmd

import (
	// "github.com/horiodino/asper/internal/pkg"
	"github.com/horiodino/asper/internal/pkg"
	"github.com/horiodino/asper/pkg"
)

type Cli struct {
	context pkg.Context
}


func checkrole() (bool, error) {

	return true, nil
}

func login() (bool, error) {
	
	return true, nil
}

func logout() (bool, error) {
	
	return true, nil
}

func isloggedIn() (bool, error) {

	return true, nil
}

func createvm(name string, vmos string, cpu string, memory string, disk string) (bool, error) {

	// check for the context
	// check all the required flags
	// check all acess control

	// then initialize  the client and call the function from the interface

	err := InitializeClient(&Cli{})
	if err != nil {
		return false, err
	}

	return true, nil
}


func InitializeClient(client *Cli) error {

	// check for the context
	value := getconxtext()

	switch value {
	case "local":
		client.context = pkg.NewLocalClient()
	case "remote":
		// client.context = pkg.NewRemoteClient()
	default:
		// return errors.New("Invalid context")
	}

	return nil
}

func getconxtext() string {
	// TODO get the context from the config file
	return "local"
}