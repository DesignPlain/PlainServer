package types

type Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirementsBaselineEbsBandwidthMbps struct {
	// The minimum baseline bandwidth, in Mbps. To specify no minimum limit, omit this parameter..
	Min int `json:"min,omitempty" yaml:"min,omitempty"`

	// The maximum baseline bandwidth, in Mbps. To specify no maximum limit, omit this parameter..
	Max int `json:"max,omitempty" yaml:"max,omitempty"`
}
