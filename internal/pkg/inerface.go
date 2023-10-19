package pkg

import "fmt"

type Context interface {
	CreateVM(name string, vmos string, cpu int, memory int, disk int) ( error)
	DeleteVM(name string) ( error)
	GetVM(name string) ( error)
	UpdateVM(name string, vmos string, cpu int, memory int, disk int) (error)
	ListVM() ( error)
	HelpVM() (error)
	VersionVM() ( error)
	ConfigVM() ( error)
	SSHVM(name string) ( error)
	MetricsVM(name string) (error)
}

type Cli struct {
	isadmin bool
}

func NewLocalClient() *Cli {
	fmt.Println("Running NewLocalClient")
	return &Cli{isadmin: true}
}