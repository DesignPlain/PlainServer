package types

type Ec2_LaunchTemplateHibernationOptions struct {
	// If set to `true`, the launched EC2 instance will hibernation enabled.
	Configured bool `json:"configured,omitempty" yaml:"configured,omitempty"`
}
