package types

type Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirementsTotalLocalStorageGb struct {
	// The maximum amount of total local storage, in GB. To specify no maximum limit, omit this parameter.
	Max float64 `json:"max,omitempty" yaml:"max,omitempty"`

	// The minimum amount of total local storage, in GB. To specify no minimum limit, omit this parameter.
	Min float64 `json:"min,omitempty" yaml:"min,omitempty"`
}
