package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "myapp",
    Short: "My CLI App",
    Long:  "A simple CLI app built with Cobra",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello from my CLI app!")
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

// cli overview 
// all commands ---> login ---> login.go
//              ---> logout ---> logout.go
//              ---> add-kvm-node ---> add-kvm-node.go   required params ---> config.go
//              ---> create ---> create.go  subcommand -----> resource type ---> resource type.go
//              ---> delete ---> delete.go  subcommand -----> resource type ---> resource type.go
//              ---> get    ---> get.go     subcommand -----> resource type ---> resource type.go
//              ---> update ---> update.go  subcommand -----> resource type ---> resource type.go
//              ---> list   ---> list.go    subcommand -----> resource type ---> resource type.go
//              ---> help   ---> help.go    subcommand -----> resource type ---> resource type.go
//              ---> version---> version.go
//              ---> config ---> config.go   ------> it will load config file and set the params in the config file 
//              ---> ssh    ---> ssh.go      ------> it will ssh to the vm
//              ---> metrics---> metrics.go  ------> it will get the metrics subcommand -----> resource type ---> resource type.go  required params ---> config.go if no name is provided it will get all the metrics of all the resorce type 



