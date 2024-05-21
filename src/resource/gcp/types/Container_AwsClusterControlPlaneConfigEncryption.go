package types

type Container_AwsClusterControlPlaneConfigEncryption struct {
	// The ARN of the AWS KMS key used to encrypt cluster configuration.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`
}
