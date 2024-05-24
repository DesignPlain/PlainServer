package types

type Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirementsTotalLocalStorageGb struct {
	// The maximum number of vCPUs. To specify no maximum limit, omit this parameter.
	Max float64 `json:"max,omitempty" yaml:"max,omitempty"`

	// The minimum number of vCPUs. To specify no minimum limit, specify `0`.
	Min float64 `json:"min,omitempty" yaml:"min,omitempty"`
}
