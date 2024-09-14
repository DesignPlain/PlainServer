package types

type Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorCount struct {
	// Minimum.
	Min int `json:"min,omitempty" yaml:"min,omitempty"`

	// Maximum. Set to `0` to exclude instance types with accelerators.
	Max int `json:"max,omitempty" yaml:"max,omitempty"`
}
