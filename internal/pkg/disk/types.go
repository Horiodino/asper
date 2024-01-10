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
	Tokens string `json:"tokens"`
}

type DiskConfigurationInput struct {
	Name string `json:"name"`
	Size int    `json:"size"`
}

type DiskConfigurationOutput struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"status"`
	Size   int    `json:"size"`
}

type DescribeDiskInput struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type DescribeDiskOutput struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	Status    string `json:"status"`
	Used_By   string `json:"used_by"`
	Size      int    `json:"size"`
	FreeSpace int    `json:"free_space"`
}

type AttachDiskInput struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type AttachDiskOutput struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"status"`
}

type DetachDiskInput struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type DetachDiskOutput struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"status"`
}

type DeleteDiskInput struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type DeleteDiskOutput struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"status"`
}
