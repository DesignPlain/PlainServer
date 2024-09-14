package types

type Ec2_LaunchTemplateInstanceMarketOptionsSpotOptions struct {
	// The maximum hourly price you're willing to pay for the Spot Instances.
	MaxPrice string `json:"maxPrice,omitempty" yaml:"maxPrice,omitempty"`

	// The Spot Instance request type. Can be `one-time`, or `persistent`.
	SpotInstanceType string `json:"spotInstanceType,omitempty" yaml:"spotInstanceType,omitempty"`

	// The end date of the request.
	ValidUntil string `json:"validUntil,omitempty" yaml:"validUntil,omitempty"`

	// The required duration in minutes. This value must be a multiple of 60.
	BlockDurationMinutes int `json:"blockDurationMinutes,omitempty" yaml:"blockDurationMinutes,omitempty"`

	/*
	   The behavior when a Spot Instance is interrupted. Can be `hibernate`,
	   `stop`, or `terminate`. (Default: `terminate`).
	*/
	InstanceInterruptionBehavior string `json:"instanceInterruptionBehavior,omitempty" yaml:"instanceInterruptionBehavior,omitempty"`
}
