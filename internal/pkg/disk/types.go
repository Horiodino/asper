package disk

import (
	"context"

	"github.com/horiodino/asper/internal/configuration"
)

type Volumes interface {
	CreateDisk(ctx context.Context, config configuration.InstanceConfigurationInput) (*configuration.InstanceConfigurationOutput, error)
	DeleteDisk(ctx context.Context, config configuration.InstanceConfigurationInput) (*configuration.InstanceConfigurationOutput, error)
	DetailsDisk(ctx context.Context, config configuration.InstanceConfigurationInput) (*configuration.InstanceConfigurationOutput, error)
	DetachDisk(ctx context.Context, config configuration.InstanceConfigurationInput) (*configuration.InstanceConfigurationOutput, error)
	AttachDisk(ctx context.Context, config configuration.InstanceConfigurationInput) (*configuration.InstanceConfigurationOutput, error)
}

type volumes struct {
	Tokens string
}
