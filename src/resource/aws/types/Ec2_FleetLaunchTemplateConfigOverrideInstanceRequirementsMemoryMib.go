package types

type Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryMib struct {
	// The minimum number of vCPUs. To specify no minimum limit, specify `0`.
	Min int `json:"min,omitempty" yaml:"min,omitempty"`

	// The maximum number of vCPUs. To specify no maximum limit, omit this parameter.
	Max int `json:"max,omitempty" yaml:"max,omitempty"`
}
