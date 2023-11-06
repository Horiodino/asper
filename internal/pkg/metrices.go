package pkg

import "fmt"

type MetricsVM interface {
	InitilizeMetriceClient() *metrice
}

type metrice struct {
	UserID     asperstring
	IsUSer     bool
	InstanceID asperstring
}

type GetMetrics interface {
	GetCPU() *metrice
	GetMemory() *metrice
	GetDisk() *metrice
	GetNetwork() *metrice
}

func (c *metrice) GetCPU() *metrice {
	fmt.Println("Running GetCPU")
	return c
}
func (c *metrice) GetMemory() *metrice {
	fmt.Println("Running GetMemory")
	return c
}
func (c *metrice) GetDisk() *metrice {
	fmt.Println("Running GetDisk")
	return c
}
func (c *metrice) GetNetwork() *metrice {
	fmt.Println("Running GetNetwork")
	return c
}

func (c *metrice) InitilizeMetriceClient() *metrice {
	fmt.Println("Running InitilizeMetriceClient")
	return c
}
