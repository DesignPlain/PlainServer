package types

type Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkBandwidthGbps struct {
	// The minimum amount of network bandwidth, in Gbps. To specify no minimum limit, omit this parameter.
	Min float64 `json:"min,omitempty" yaml:"min,omitempty"`

	// The maximum amount of network bandwidth, in Gbps. To specify no maximum limit, omit this parameter.
	Max float64 `json:"max,omitempty" yaml:"max,omitempty"`
}
