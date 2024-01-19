package vm

import "fmt"

type Describe_VM interface {
	DescribeVM()
}

func (c *vm) DescribeVM() {
	fmt.Println("Running DescribeVM")
}

//
//import (
//	"fmt"
//	"github.com/horiodino/asper/internal/logger"
//)
//
//type list struct {
//	UserID logger.Asperstring
//	IsUSer bool
//}
//
//type ListOfResources interface {
//	InitilizeListClient() *list
//}
//
//type List interface {
//	ListVM() *UserContext
//	ListFirewall() *UserContext
//	ListNetwork() *UserContext
//	ListDisk() *UserContext
//	ListSSHKey() *UserContext
//}
//
//func (c *list) InitilizeListClient() *list {
//	fmt.Println("Running InitilizeListClient")
//	return c
//}
//
//func (c *list) ListVM() *list {
//	fmt.Println("Running ListVM")
//	return c
//}
//func (c *list) ListFirewall() *list {
//	fmt.Println("Running ListFirewall")
//	return c
//}
//func (c *list) ListNetwork() *list {
//	fmt.Println("Running ListNetwork")
//	return c
//}
//func (c *list) ListDisk() *list {
//	fmt.Println("Running ListDisk")
//	return c
//}
//func (c *list) ListSSHKey() *list {
//	fmt.Println("Running ListSSHKey")
//	return c
//}
