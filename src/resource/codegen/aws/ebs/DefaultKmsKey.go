package ebs

type DefaultKmsKey struct {
	// The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use to encrypt the EBS volume.
	KeyArn string `json:"keyArn,omitempty" yaml:"keyArn,omitempty"`
}
