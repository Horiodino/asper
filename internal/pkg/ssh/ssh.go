package ssh

import "context"

<<<<<<< HEAD
func CreateSSH(ctx context.Context, input SSHConfigurationInput) (SSHKeyOutput, error) {
	return SSHKeyOutput{}, nil
}

func DeleteSSH(ctx context.Context, input DescribeSSHKeyInput) (DescribeSSHKeyOutput, error) {
	return DescribeSSHKeyOutput{}, nil
}

func DescribeSSH(ctx context.Context, input DescribeSSHKeyInput) (DescribeSSHKeyOutput, error) {
	return DescribeSSHKeyOutput{}, nil
=======
func CreateSSH(ctx context.Context,  InputSSHConfiguration) (InstanceConfigurationOutput, error) {
	return InstanceConfigurationOutput{}, nil
}

func DeleteSSH(ctx context.Context,  InputSSHConfiguration) (InstanceConfigurationOutput, error) {
	return InstanceConfigurationOutput{}, nil
}

func DescribeSSH(ctx context.Context,  InputSSHConfiguration) (InstanceConfigurationOutput, error) {
	return InstanceConfigurationOutput{}, nil
>>>>>>> 41416dd19ca273ab72348318ae3e611423cdb77c
}
