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

var createvms = &cobra.Command{
    Use:   "createvm",
    Short: " ",
    Long:  "A simple CLI app built with Cobra",
    Run: func(cmd *cobra.Command, args []string) {


        // TODO  check the context , in which context the cli is running is it local or remote 
        // if its remote then we will skip the authentication from database and we will use the token from the config file
        // and if remote then we will fetch from database and we will use the token from the config file
        // it will be done in the config.go file

        respone, err := isloggedIn()
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        if respone == false {
            fmt.Println("Please login first")
            os.Exit(1)
        }

        respone, err = checkrole()  // foe rbac
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        if respone == false {
            fmt.Println("You don't have permission to create vm")
            os.Exit(1)
        }


        // USE INTERFACE WE WILL MAKE TWO CLIENTS ONE FOR LOCAL AND ONE FOR REMOTE
        // in side all use interface and we will call the function from the interface
        respone, err = createvm(name, vmos, cpu, memory, disk)


    },
}

var context = &cobra.Command{
    Use:   "context",
    Short: " set the context of the cli local or remote",
}

var use = &cobra.Command{
    Use : "set the context ",
    Short : "set the context of the cli local or remote",
    Run : func(cmd *cobra.Command, args []string) {
        // TODO set the context of the cli local or remote
        // if local then we will use the token from the config file
        // if remote then we will fetch from database and we will use the token from the config file
        // it will be done in the config.go file
    },
}


func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

// add subcommand for context

func init () {
    rootCmd.AddCommand(createvms)
    rootCmd.AddCommand(context)
    context.AddCommand(use)
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



