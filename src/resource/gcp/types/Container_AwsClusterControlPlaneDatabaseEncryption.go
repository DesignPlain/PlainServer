package types

type Container_AwsClusterControlPlaneDatabaseEncryption struct {
	// The ARN of the AWS KMS key used to encrypt cluster secrets.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`
}
