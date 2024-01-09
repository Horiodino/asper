package qemu

import (
	"fmt"
	"log"

	libvertgo "github.com/libvirt/libvirt-go"
)

type QEMU struct {
	ResourceName string
	Client       *libvertgo.Connect
}

func qemuclient() *libvertgo.Connect {
	conn, err := libvertgo.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("Error connecting to hypervisor:", err)
		log.Fatal(err)
	}

	return conn
}

func InitializeQEMU() *QEMU {
	return &QEMU{
		ResourceName: "nil",
		Client:       qemuclient(),
	}
}
