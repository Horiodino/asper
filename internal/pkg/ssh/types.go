package ssh

type Ssh interface {
	CreateSSH(ctx context.Context, config configuration.InstanceConfigurationInput) (*configuration.InstanceConfigurationOutput, error)
}