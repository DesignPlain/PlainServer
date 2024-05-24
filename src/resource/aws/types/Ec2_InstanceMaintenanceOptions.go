package types

type Ec2_InstanceMaintenanceOptions struct {
	// Automatic recovery behavior of the Instance. Can be `"default"` or `"disabled"`. See [Recover your instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-recover.html) for more details.
	AutoRecovery string `json:"autoRecovery,omitempty" yaml:"autoRecovery,omitempty"`
}
