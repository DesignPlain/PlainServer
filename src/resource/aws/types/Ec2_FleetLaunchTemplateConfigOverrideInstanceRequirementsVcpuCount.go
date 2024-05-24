package types

type Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirementsVcpuCount struct {
	// The maximum number of vCPUs. To specify no maximum limit, omit this parameter.
	Max int `json:"max,omitempty" yaml:"max,omitempty"`

	// The minimum number of vCPUs. To specify no minimum limit, specify `0`.
	Min int `json:"min,omitempty" yaml:"min,omitempty"`
}
