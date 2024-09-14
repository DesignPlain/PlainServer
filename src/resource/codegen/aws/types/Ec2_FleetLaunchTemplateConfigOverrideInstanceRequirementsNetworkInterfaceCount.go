package types

type Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkInterfaceCount struct {
	// The maximum number of network interfaces. To specify no maximum limit, omit this parameter.
	Max int `json:"max,omitempty" yaml:"max,omitempty"`

	// The minimum number of network interfaces. To specify no minimum limit, omit this parameter.
	Min int `json:"min,omitempty" yaml:"min,omitempty"`
}
