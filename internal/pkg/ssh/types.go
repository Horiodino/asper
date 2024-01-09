package ssh

import (
	"context"
)

type Ssh interface {
	CreateSSH(ctx context.Context, config InputSSHConfiguration) (InstanceConfigurationOutput, error)
	DeleteSSH(ctx context.Context, config InputSSHConfiguration) (InstanceConfigurationOutput, error)
	DescribeSSH(ctx context.Context, config InputSSHConfiguration) (InstanceConfigurationOutput, error)
}

type ssh struct {
	Token string `json:"token"`
}

func NewSSHClient() *ssh {
	return &ssh{}
}
