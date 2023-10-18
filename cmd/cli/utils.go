package cmd

import (
	"fmt"

	"github.com/horiodino/asper/internal/pkg"
)

type Cli struct {
	context pkg.Context
}


func checkrole() (bool, error) {
fmt.Println("Running checkrole")
	return true, nil
}

func login() (bool, error) {
	fmt.Println("Running login")
	return true, nil
}

func logout() (bool, error) {
	fmt.Println("Running logout")
	return true, nil
}

func isloggedIn() (bool, error) {
	fmt.Println("Running isloggedIn")
	return true, nil
}

func createvm(name string, vmos string, cpu string, memory string, disk string) (bool, error) {

	// check for the context
	// check all the required flags
	// check all acess control
		fmt.Println("Running createvm")
	// then initialize  the client and call the function from the interface

	client , err := InitializeClient(&Cli{})
	if err != nil {
		return false, err
	}

	fmt.Println("Successfully created vm")

	client.context.CreateVM(name, vmos, 1,2, 10)


	return true, nil
}


func InitializeClient(client *Cli) (*Cli, error) {

	fmt.Println("Running InitializeClient")

	// check for the context
	value := getconxtext()

	switch value {
	case "local":
		client.context = pkg.NewLocalClient()

		return client, nil

	case "remote":
		// client.context = pkg.NewRemoteClient()
	default:
		// return errors.New("Invalid context")
	}

	return nil, nil
}

func getconxtext() string {
	// TODO get the context from the config file
	return "local"
}