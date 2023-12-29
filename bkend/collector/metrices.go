package vm

import (
	"fmt"

	"github.com/horiodino/asper/internal/logger"
)

// TODO add use valid  argument for all methods

type MetricsVM interface {
	InitilizeMetriceClient() *metrice
}

type metrice struct {
	UserID     logger.Asperstring
	IsUSer     bool
	InstanceID logger.Asperstring
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
