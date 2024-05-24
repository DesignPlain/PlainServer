package types

type Ec2_SpotFleetRequestLaunchTemplateConfig struct {
	// Launch template specification. See Launch Template Specification below for more details.
	LaunchTemplateSpecification Ec2_SpotFleetRequestLaunchTemplateConfigLaunchTemplateSpecification `json:"launchTemplateSpecification,omitempty" yaml:"launchTemplateSpecification,omitempty"`

	// One or more override configurations. See Overrides below for more details.
	Overrides []Ec2_SpotFleetRequestLaunchTemplateConfigOverride `json:"overrides,omitempty" yaml:"overrides,omitempty"`
}
