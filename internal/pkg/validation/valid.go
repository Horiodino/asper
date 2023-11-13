package validation_test

import (
	"fmt"

	"github.com/horiodino/asper/internal/auth"
	"github.com/horiodino/asper/internal/pkg"
)

var authclient, _ = auth.InitializeAuth()

// Validator is the interface for validation operations.
type Validator interface {
	CheckRole() (bool, error)
	Login() (bool, error)
	Logout() (bool, error)
	IsLoggedIn() (bool, error)
	InitilizeCreateClient() *pkg.Create
}

// CliValidator implements the Validator interface for CLI operations.
type CliValidator struct {
	context pkg.Asperstring
}

func (c *CliValidator) CheckRole() (bool, error) {

}

func (c *CliValidator) Login() (bool, error) {
	switch c.context {
	case "local":
		switch c.context {
		case "local":
			val, err := authclient.Login()
			if err != nil {
				panic(err)
			}
			if val == false {
				fmt.Println("NOOOOOOOOOOOOOOOOOOO")
			}
			fmt.Println("Running local checkrole------------------------------")
			// check in .asper
		case "cloud":
			fmt.Println("Running cloud checkrole")
			// check in database
		default:
			fmt.Println("Running default checkrole")
		}
		fmt.Println("Running checkrole")
		return true, nil
		// check in .asper
	case "cloud":
		fmt.Println("Running cloud login")
		// check in database
	default:
		fmt.Println("Running default login")
	}
	fmt.Println("Running login")
	return true, nil
}

func (c *CliValidator) Logout() (bool, error) {
	switch c.context {
	case "local":
		fmt.Println("Running local logout")
		// check in .asper
	case "cloud":
		fmt.Println("Running cloud logout")
		// check in database
	default:
		fmt.Println("Running default logout")
	}
	fmt.Println("Running logout")
	return true, nil
}

func (c *CliValidator) IsLoggedIn() (bool, error) {
	switch c.context {
	case "local":
		fmt.Println("Running local isloggedIn")
		// check in .asper
	case "cloud":
		fmt.Println("Running cloud isloggedIn")
		// check in database
	default:
		fmt.Println("Running default isloggedIn")
	}
	fmt.Println("Running isloggedIn")
	return true, nil
}

// InitializeClient initializes the client based on the context.
func InitializeClient() (*CliValidator, error) {
	fmt.Println("Running InitializeClient")

	// Check for the context
	value := getContext()

	switch value {
	case "local":
		return &CliValidator{context: "local"}, nil
	case "remote":
		// client.(*CliValidator).context = pkg.NewRemoteClient()
	default:
		// return nil, errors.New("Invalid context")
		return nil, nil
	}

	return nil, nil
}

// getContext gets the context from the config file.
func getContext() string {
	// TODO: Get the context from the config file
	return "local"
}
