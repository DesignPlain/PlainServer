package types

type Ec2_getLaunchTemplateInstanceMarketOptionSpotOption struct {
	//
	MaxPrice string `json:"maxPrice,omitempty" yaml:"maxPrice,omitempty"`

	//
	SpotInstanceType string `json:"spotInstanceType,omitempty" yaml:"spotInstanceType,omitempty"`

	//
	ValidUntil string `json:"validUntil,omitempty" yaml:"validUntil,omitempty"`

	//
	BlockDurationMinutes int `json:"blockDurationMinutes,omitempty" yaml:"blockDurationMinutes,omitempty"`

	//
	InstanceInterruptionBehavior string `json:"instanceInterruptionBehavior,omitempty" yaml:"instanceInterruptionBehavior,omitempty"`
}
