package types

type Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorTotalMemoryMib struct {
	// The maximum amount of accelerator memory, in MiB. To specify no maximum limit, omit this parameter.
	Max int `json:"max,omitempty" yaml:"max,omitempty"`

	// The minimum amount of accelerator memory, in MiB. To specify no minimum limit, omit this parameter.
	Min int `json:"min,omitempty" yaml:"min,omitempty"`
}
