package types

type Securitylake_CustomLogSourceProviderDetail struct {
	// The location of the partition in the Amazon S3 bucket for Security Lake.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role to be used by the AWS Glue crawler.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
