package cloudinit

import "context"

type CloudInit interface {
	CreateCloudInit(ctx context.Context, config *CloudInitInput) (*CloudInitOutput, error)
	DeleteCloudInit(ctx context.Context, config *CloudInitInput) (*CloudInitOutput, error)
}

type CloudInitInput struct {
	SSHConfiguration InputSSHConfiguration
	Content          string
}

type CloudInitOutput struct {
	ID         string
	ResultMeta string
}
