package disk

import (
	"context"
)

const (
	ALLOCATED = "allocated"
	ATTACHED  = "attached"
	DETACHED  = "detached"
	DELETED   = "deleted"
)

type Volumes interface {
	CreateDisk(ctx context.Context, config DiskConfigurationInput) (*DiskConfigurationOutput, error)
	DeleteDisk(ctx context.Context, config DeleteDiskInput) (*DeleteDiskOutput, error)
	DetailsDisk(ctx context.Context, config DescribeDiskInput) (*DescribeDiskOutput, error)
	DetachDisk(ctx context.Context, config DetachDiskInput) (*DetachDiskOutput, error)
	AttachDisk(ctx context.Context, config AttachDiskInput) (*AttachDiskOutput, error)
}

type volumes struct {
	Tokens string
}

type DiskConfigurationInput struct {
	Name string
	Size int
}

type DiskConfigurationOutput struct {
	Name   string
	ID     string
	Status string
	Size   int
}

type DescribeDiskInput struct {
	Name string
	ID   string
}

type DescribeDiskOutput struct {
	Name      string
	ID        string
	Status    string
	Used_By   string
	Size      int
	FreeSpace int
}

type AttachDiskInput struct {
	Name string
	ID   string
}

type AttachDiskOutput struct {
	Name   string
	ID     string
	Status string
}

type DetachDiskInput struct {
	Name string
	ID   string
}

type DetachDiskOutput struct {
	Name   string
	ID     string
	Status string
}

type DeleteDiskInput struct {
	Name string
	ID   string
}

type DeleteDiskOutput struct {
	Name   string
	ID     string
	Status string
}
