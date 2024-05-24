package types

type Ec2_FleetLaunchTemplateConfig struct {
	// Nested argument containing EC2 Launch Template to use. Defined below.
	LaunchTemplateSpecification Ec2_FleetLaunchTemplateConfigLaunchTemplateSpecification `json:"launchTemplateSpecification,omitempty" yaml:"launchTemplateSpecification,omitempty"`

	// Nested argument(s) containing parameters to override the same parameters in the Launch Template. Defined below.
	Overrides []Ec2_FleetLaunchTemplateConfigOverride `json:"overrides,omitempty" yaml:"overrides,omitempty"`
}
