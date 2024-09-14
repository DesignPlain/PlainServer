package types

type Ec2_LaunchTemplateMonitoring struct {
	// If `true`, the launched EC2 instance will have detailed monitoring enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
