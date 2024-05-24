package types

type Dlm_LifecyclePolicyPolicyDetailsActionCrossRegionCopyEncryptionConfiguration struct {
	// The Amazon Resource Name (ARN) of the AWS KMS customer master key (CMK) to use for EBS encryption. If this argument is not specified, the default KMS key for the account is used.
	CmkArn string `json:"cmkArn,omitempty" yaml:"cmkArn,omitempty"`

	// To encrypt a copy of an unencrypted snapshot if encryption by default is not enabled, enable encryption using this parameter. Copies of encrypted snapshots are encrypted, even if this parameter is false or if encryption by default is not enabled.
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`
}
