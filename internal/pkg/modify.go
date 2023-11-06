package pkg

import (
	"fmt"
)

type Modify interface {
	InitilizeModifyClient(instanceid asperstring) *MmodifyInstance
}

type MmodifyInstance struct {
	UserID asperstring
	IsUSer bool
}

func (c *MmodifyInstance) InitilizeModifyClient(instanceid asperstring) *MmodifyInstance {
	fmt.Println("Running InitilizeModifyClient")
	return c
}

type modifyclient struct {
	UserID asperstring
	IsUSer bool
}

type ModifyInstance interface {
	ModifyVM(instanceid asperstring) *MmodifyInstance
}

func (c *modifyclient, p *MmodifyInstance) ModifyVM(instanceid asperstring) *modifyclient {
	fmt.Println("Running ModifyVM")
	return c
}

type ModifyOperations interface {
	MoveVM(instanceid asperstring) *modifyclient
	ResizeVM(instanceid asperstring) *modifyclient
	RebootVM(instanceid asperstring) *modifyclient
	StopVM(instanceid asperstring) *modifyclient
	StartVM(instanceid asperstring) *modifyclient
	RebuildVM(instanceid asperstring) *modifyclient
	ReinstallVM(instanceid asperstring) *modifyclient
	ResetVM(instanceid asperstring) *modifyclient
	CloneVM(instanceid asperstring) *modifyclient
	BackupVM(instanceid asperstring) *modifyclient
	RestoreVM(instanceid asperstring) *modifyclient
	ConvertVM(instanceid asperstring) *modifyclient
	ModifyFirewall(instanceid asperstring) *modifyclient
	ModifyNetwork(instanceid asperstring) *modifyclient
	ModifyDisk(instanceid asperstring) *modifyclient
	ModifySSH(instanceid asperstring) *modifyclient
}

/*
	// note we will make the modify instance type such that it can be used for all the modify operations
	// we will use the same type for all the modify operations

	how we will do it?
	we will create a type ModifyInstance and then we will create a method for each operation
	we will then use the same type for all the modify operations
*/

func (c *modifyclient) MoveVM(instanceid asperstring) *modifyclient {
	fmt.Println("Running MoveVM")
	return c
}
func (c *modifyclient) ResizeVM(instanceid asperstring) *modifyclient {
	fmt.Println("Running ResizeVM")
	return c
}
func (c *modifyclient) RebootVM(instanceid asperstring) *modifyclient {
	fmt.Println("Running RebootVM")
	return c
}
func (c *modifyclient) StopVM(instanceid asperstring) *modifyclient {
	fmt.Println("Running StopVM")
	return c
}
func (c *modifyclient) StartVM(instanceid asperstring) *modifyclient {
	fmt.Println("Running StartVM")
	return c
}
func (c *modifyclient) RebuildVM(instanceid asperstring) *modifyclient {
	fmt.Println("Running RebuildVM")
	return c
}
func (c *modifyclient) ReinstallVM(instanceid asperstring) *modifyclient {
	fmt.Println("Running ReinstallVM")
	return c
}
func (c *modifyclient) ResetVM(instanceid asperstring) *modifyclient {
	fmt.Println("Running ResetVM")
	return c
}
func (c *modifyclient) CloneVM(instanceid asperstring) *modifyclient {
	fmt.Println("Running CloneVM")
	return c
}
func (c *modifyclient) BackupVM(instanceid asperstring) *modifyclient {
	fmt.Println("Running BackupVM")
	return c
}
func (c *modifyclient) RestoreVM(instanceid asperstring) *modifyclient {
	fmt.Println("Running RestoreVM")
	return c
}
func (c *modifyclient) ConvertVM(instanceid asperstring) *modifyclient {
	fmt.Println("Running ConvertVM")
	return c
}
func (c *modifyclient) modifyclientModifyFirewall(instanceid asperstring) *modifyclient {
	fmt.Println("Running ModifyFirewall")
	return c
}
func (c *modifyclient) ModifyNetwork(instanceid asperstring) *modifyclient {
	fmt.Println("Running ModifyNetwork")
	return c
}
func (c *modifyclient) ModifyDisk(instanceid asperstring) *modifyclient {
	fmt.Println("Running ModifyDisk")
	return c
}
func (c *modifyclient) ModifySSH(instanceid asperstring) *modifyclient {
	fmt.Println("Running ModifySSH")
	return c
}
