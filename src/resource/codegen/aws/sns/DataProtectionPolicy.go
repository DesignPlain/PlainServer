package sns

type DataProtectionPolicy struct {
	// The ARN of the SNS topic
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// The fully-formed AWS policy as JSON. For more information about building AWS IAM policy documents with this provider, see the AWS IAM Policy Document Guide.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
