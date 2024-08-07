package types

type Athena_WorkgroupConfigurationResultConfigurationEncryptionConfiguration struct {
	// Whether Amazon S3 server-side encryption with Amazon S3-managed keys (`SSE_S3`), server-side encryption with KMS-managed keys (`SSE_KMS`), or client-side encryption with KMS-managed keys (`CSE_KMS`) is used. If a query runs in a workgroup and the workgroup overrides client-side settings, then the workgroup's setting for encryption is used. It specifies whether query results must be encrypted, for all queries that run in this workgroup.
	EncryptionOption string `json:"encryptionOption,omitempty" yaml:"encryptionOption,omitempty"`

	// For `SSE_KMS` and `CSE_KMS`, this is the KMS key ARN.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`
}
