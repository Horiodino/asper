package auth

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

var configFilePath = "~/.asper/config.yaml"
var FilePath, _ = homedir.Expand(configFilePath)

// this will implement all the auth related function like login, logout, checkrole, etc

type Auth interface {
	Login() (bool, error)
	Logout() (bool, error)
	IsLoggedIn() (bool, error)
	CheckRole() (bool, error)
}

// TODO we will authenticate the user from the database in case of the remote context
// we will create the db for the first time and we will store the token in the config file
// once the db is created we will use the token from the config file and we will authenticate the user

// for the cli context point of view we will use the token from the config file and we will authenticate the user
// but we dont use any database for the cli context we use simple .config file for the cli context to save the token
// and all other information

// -------------------------------------------------------------------------------------------

type Config struct {
	Username  string    `yaml:"username"`
	Password  string    `yaml:"password"`
	LastLogin time.Time `yaml:"last_login"`
}

// Login checks .asper/config file for the role/permission and performs authentication.
func (auth *Config) Login(filePath string) (bool, error) {
	// Read the configuration file
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return false, err
	}

	// Unmarshal the configuration
	var config Config
	err = yaml.Unmarshal(fileContent, &config)
	if err != nil {
		return false, err
	}

	// Authenticate user
	if auth.Username == config.Username && auth.Password == config.Password {
		fmt.Println(config.Username)
		fmt.Println(config.Password)

		// Check if the session has expired
		if !IsSessionExpired(config.LastLogin) {
			fmt.Println("User session is still active.")
			return true, nil
		}

		fmt.Println("User session expired plese IMPLEMENT ME ")
		return true, nil
	}

	return false, fmt.Errorf("Invalid user")
}

// IsSessionExpired checks if the user's session has expired based on the last login timestamp.
func IsSessionExpired(lastLogin time.Time) bool {
	sessionTimeout := 12 * time.Hour // Adjust as needed

	return time.Since(lastLogin) > sessionTimeout
}

//-----------------------------------------------------------------------------------------------------------------

func (auth *Config) Logout() (bool, error) {

	return true, nil
}

func (auth *Config) IsLoggedIn() (bool, error) {

	return true, nil
}

func (auth *Config) CheckRole() (bool, error) {

	return true, nil
}

func InitializeAuth() (*Config, error) {
	return &Config{
		Username:  "your_username",
		Password:  "your_password",
		LastLogin: time.Now(),
	}, nil
}
