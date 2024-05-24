package s3control

type AccessPointPolicy struct {
	// The ARN of the access point that you want to associate with the specified policy.
	AccessPointArn string `json:"accessPointArn,omitempty" yaml:"accessPointArn,omitempty"`

	// The policy that you want to apply to the specified access point.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
