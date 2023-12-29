package vm

import (
	"context"
	"fmt"
	"time"

	"github.com/horiodino/asper/internal/configuration"
	"github.com/horiodino/asper/internal/logger"
	"github.com/libvirt/libvirt-go"
)

// TODO add use valid  argument for all methods

type vm struct {
	Token string `json:"token"` // Authentication token for the VM operations
}

func InitializeFromConfig() *vm {

	return &vm{}
}

// Create interface groups together methods for creating various components of a VM
type Create interface {
	CreateVM(ctx context.Context, input *configuration.InstanceConfigurationInput) (*configuration.InstanceConfigurationOutput, error)
	DeleteVM(ctx context.Context)
}

func (c *vm) CreateVM(ctx context.Context, input *configuration.InstanceConfigurationInput, operationType string) (*configuration.InstanceConfigurationOutput, error) {

	fmt.Println("working!!!!")
	switch operationType {
	case "managed":
	case "custom":
	default:
	}

	// Creation of VM
	logger := logger.NewLogger()
	logger.LogInfo("Running CreateVM")

	time.Sleep(15 * time.Second)
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		logger.LogError(fmt.Sprintf("Error connecting to hypervisor: %v", err))
		return nil, err
	}
	defer conn.Close()

	domain, err := conn.DomainDefineXML(input)
	if err != nil {
		logger.LogError(fmt.Sprintf("Error defining domain: %v", err))
		return nil, err
	}
	defer domain.Free()

	logger.LogInfo("Initializing virtual machine components")

	err = domain.Create()
	if err != nil {
		logger.LogError(fmt.Sprintf("Error starting domain: %v", err))
		return nil, err
	}

	logger.LogInfo("Virtual machine started successfully.")
	logger.LogInfo("Initializing virtual machine components. Please wait for a few seconds.")

	return nil, nil
}
