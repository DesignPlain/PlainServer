package types

type Ec2_SpotFleetRequestLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorCount struct {
	// Maximum. Set to `0` to exclude instance types with accelerators.
	Max int `json:"max,omitempty" yaml:"max,omitempty"`

	// Minimum.
	Min int `json:"min,omitempty" yaml:"min,omitempty"`
}
