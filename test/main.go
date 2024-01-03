package main

import (
	"context"
	"log"

	//	"github.com/horiodino/asper/internal/pkg/network"
	"github.com/horiodino/asper/internal/pkg/vm"
)
func main() {
	ctx:= context.Background()
//	Nclient := network.NewNetworkClient()
//	_, err := Nclient.CreateNetworkInterface(ctx,&network.NetworkInterfaceParams{
//		Name: "testing-nic",
//	})
//	if err != nil {
//		log.Fatal(err)
//	}

	Client := vm.InitializeFromConfig()
	_, err := Client.CreateVM(ctx,&vm.InstanceConfigurationInput{})
	if err!= nil{
		log.Fatal(err)
	}

}
