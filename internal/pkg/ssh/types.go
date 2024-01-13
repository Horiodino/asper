package ssh

import (
	"context"

	"github.com/horiodino/asper/internal/logger"
)

type Ssh interface {
	CreateSSH(ctx context.Context, config SSHConfigurationInput) (SSHKeyOutput, error)
	DeleteSSH(ctx context.Context, config DescribeSSHKeyInput) (DescribeSSHKeyOutput, error)
	DescribeSSH(ctx context.Context, config DescribeSSHKeyInput) (DescribeSSHKeyOutput, error)
}

type SSHConfigurationInput struct {
	Keyname string `json:"keyname"`
}
type SSHKeyOutput struct {
	Keyname string `json:"keyname"`
}
type DescribeSSHKeyInput struct {
	Name string `json:"name"`
	// only any one field is valid
	InstanceID string `json:"instance_id"`
}

type DescribeSSHKeyOutput struct {
	Keyname string               `json:"keyname"`
	UsedBy  []logger.Asperstring `json:"used_by"`
}

type ssh struct {
	Token string `json:"token"`
}

func NewSSHClient() *ssh {
	return &ssh{}
}
