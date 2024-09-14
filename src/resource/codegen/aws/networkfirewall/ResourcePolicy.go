package networkfirewall

type ResourcePolicy struct {
	// JSON formatted policy document that controls access to the Network Firewall resource. The policy must be provided --without whitespaces--.  We recommend using jsonencode for formatting as seen in the examples above. For more details, including available policy statement Actions, see the [Policy](https://docs.aws.amazon.com/network-firewall/latest/APIReference/API_PutResourcePolicy.html#API_PutResourcePolicy_RequestSyntax) parameter in the AWS API documentation.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// The Amazon Resource Name (ARN) of the rule group or firewall policy.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`
}
