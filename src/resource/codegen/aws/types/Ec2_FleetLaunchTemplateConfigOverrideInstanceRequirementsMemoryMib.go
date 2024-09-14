package types

type Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryMib struct {
	// The maximum amount of memory, in MiB. To specify no maximum limit, omit this parameter.
	Max int `json:"max,omitempty" yaml:"max,omitempty"`

	// The minimum amount of memory, in MiB. To specify no minimum limit, specify `0`.
	Min int `json:"min,omitempty" yaml:"min,omitempty"`
}
