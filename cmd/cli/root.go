package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)
var (
    name string
    vmos string
    cpu string
    memory string
    disk string
)

var rootCmd = &cobra.Command{
    Use:   "asper",
    Short: " ",
    Long:  "A simple CLI app built with Cobra",
}

var createvm = &cobra.Command{
    Use:   "createvm",
    Short: " ",
    Long:  "A simple CLI app built with Cobra",
    Run: func(cmd *cobra.Command, args []string) {

        Cli := pkg.Cli{}

        respone, err := Cli.IsLoggedIn()
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        if respone == false {
            fmt.Println("Please login first")
            os.Exit(1)
        }

        respone, err = Cli.checkrole()
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        if respone == false {
            fmt.Println("You don't have permission to create vm")
            os.Exit(1)
        }

        respone, err = Cli.createvm(name, vmos, cpu, memory, disk)


    },
}


func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init () {
    rootCmd.AddCommand(createvm)
    requiredflags()
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



