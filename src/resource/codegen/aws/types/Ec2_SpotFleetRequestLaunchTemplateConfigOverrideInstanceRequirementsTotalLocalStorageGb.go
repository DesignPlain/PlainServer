package types

type Ec2_SpotFleetRequestLaunchTemplateConfigOverrideInstanceRequirementsTotalLocalStorageGb struct {
	// Maximum. May be a decimal number, e.g. `0.5`.
	Max float64 `json:"max,omitempty" yaml:"max,omitempty"`

	// Minimum. May be a decimal number, e.g. `0.5`.
	Min float64 `json:"min,omitempty" yaml:"min,omitempty"`
}
