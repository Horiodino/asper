/*

package main

import (
	validate "github.com/horiodino/asper/internal/pkg/validation"
)

func main() {

	// before calling the main execution provess we create a new pkg that will validate every
	// request and validate the input and took security measures.

	validate, err := validate.InitializeClient()
	if err != nil {
		panic(err)
	}
	validate.Login()
	validate.Create("vm")

	// asper := test.NewLocalClient()
	// metricesclient := asper.MetricsVM.InitilizeMetriceClient()
	// metricesclient.GetCPU()
}




*/

package main

import (
	"context"
	"fmt"
	"github.com/horiodino/asper/internal/configuration"
	"github.com/horiodino/asper/internal/pkg/vm"
	"log"
)

func main() {
	// Initialize vm
	vmclient := vm.InitializeFromConfig()
	responce, err := vmclient.CreateVM(context.TODO(), &configuration.InstanceConfigurationInput{}, "")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(responce)

}
