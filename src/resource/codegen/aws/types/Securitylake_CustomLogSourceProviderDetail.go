package types

type Securitylake_CustomLogSourceProviderDetail struct {
	// The location of the partition in the Amazon S3 bucket for Security Lake.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The ARN of the IAM role to be used by the entity putting logs into your custom source partition.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
