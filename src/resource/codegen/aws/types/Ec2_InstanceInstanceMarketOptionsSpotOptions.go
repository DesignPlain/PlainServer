package types

type Ec2_InstanceInstanceMarketOptionsSpotOptions struct {
	// The behavior when a Spot Instance is interrupted. Valid values include `hibernate`, `stop`, `terminate` . The default is `terminate`.
	InstanceInterruptionBehavior string `json:"instanceInterruptionBehavior,omitempty" yaml:"instanceInterruptionBehavior,omitempty"`

	// The maximum hourly price that you're willing to pay for a Spot Instance.
	MaxPrice string `json:"maxPrice,omitempty" yaml:"maxPrice,omitempty"`

	// The Spot Instance request type. Valid values include `one-time`, `persistent`. Persistent Spot Instance requests are only supported when the instance interruption behavior is either hibernate or stop. The default is `one-time`.
	SpotInstanceType string `json:"spotInstanceType,omitempty" yaml:"spotInstanceType,omitempty"`

	// The end date of the request, in UTC format (YYYY-MM-DDTHH:MM:SSZ). Supported only for persistent requests.
	ValidUntil string `json:"validUntil,omitempty" yaml:"validUntil,omitempty"`
}
