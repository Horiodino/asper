package validation_test

import (
	"fmt"

	"github.com/horiodino/asper/internal/auth"
	"github.com/horiodino/asper/internal/configuration"

	// "github.com/horiodino/asper/internal/pkg"
	logger "github.com/horiodino/asper/internal/logger"
)

var authclient, _ = auth.InitializeAuth()

// Validator is the interface for validation operations.
type Validator interface {
	CheckRole() (bool, error)
	Login() (bool, error)
	Logout() (bool, error)
	IsLoggedIn() (bool, error)
	// InitilizeCreateClient() *pkg.Create
}

// CliValidator implements the Validator interface for CLI operations.
type CliValidator struct {
	context logger.Asperstring
	request Request
}

func (c *CliValidator) CheckRole() (bool, error) {

	return true, nil
}

func (c *CliValidator) Login() (bool, error) {
	switch c.context {
	case "local":
		switch c.context {
		case "local":
			val := true
			if val == false {
				fmt.Println("NOOOOOOOOOOOOOOOOOOO")
			}
			logger := logger.NewLogger()
			logger.LogInfo("Running local login")

			return true, nil
		case "cloud":
			fmt.Println("Running cloud checkrole")
		default:
			fmt.Println("Running default checkrole")
		}
		return true, nil
	case "cloud":
		fmt.Println("Running cloud login")
	default:
		fmt.Println("Running default login")
	}

	return false, nil
}

func (c *CliValidator) Logout() (bool, error) {
	switch c.context {
	case "local":
		fmt.Println("Running local logout")
	case "cloud":
		fmt.Println("Running cloud logout")
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
	case "cloud":
		fmt.Println("Running cloud isloggedIn")
	default:
		fmt.Println("Running default isloggedIn")
	}
	fmt.Println("Running isloggedIn")
	return true, nil
}

// InitializeClient initializes the client based on the context.
func InitializeClient() (*CliValidator, error) {

	logs := logger.NewLogger()
	logs.LogInfo("Running InitializeClient")

	// Check for the context
	value := getContext()

	switch value {
	case "local":
		return &CliValidator{
			context: "local",
			request: Request{
				Token:       "token",
				RequestType: "managed",
				Service:     "vm",
				Configuration: &configuration.ALL{
					InstanceConfigurationInput: configuration.InstanceConfigurationInput{
						Name:   "testinstance",
						OS:     "ubuntu",
						Memory: 1097152,
						VCPU:   3,
					},
				},
			},
		}, nil
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
