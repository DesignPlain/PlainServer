package types

type Ec2_VpcIpamPoolCidrCidrAuthorizationContext struct {
	// The plain-text authorization message for the prefix and account.
	Message string `json:"message,omitempty" yaml:"message,omitempty"`

	// The signed authorization message for the prefix and account.
	Signature string `json:"signature,omitempty" yaml:"signature,omitempty"`
}
