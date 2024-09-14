package types

type Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryGibPerVcpu struct {
	// The maximum amount of memory per vCPU, in GiB. To specify no maximum limit, omit this parameter.
	Max float64 `json:"max,omitempty" yaml:"max,omitempty"`

	// The minimum amount of memory per vCPU, in GiB. To specify no minimum limit, omit this parameter.
	Min float64 `json:"min,omitempty" yaml:"min,omitempty"`
}
