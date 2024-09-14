package types

type Emr_BlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRange struct {
	// The final port in the range of TCP ports.
	MaxRange int `json:"maxRange,omitempty" yaml:"maxRange,omitempty"`

	// The first port in the range of TCP ports.
	MinRange int `json:"minRange,omitempty" yaml:"minRange,omitempty"`
}
