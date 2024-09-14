package timestreamwrite

type Database struct {
	// The name of the Timestream database. Minimum length of 3. Maximum length of 64.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// The ARN (not Alias ARN) of the KMS key to be used to encrypt the data stored in the database. If the KMS key is not specified, the database will be encrypted with a Timestream managed KMS key located in your account. Refer to [AWS managed KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk) for more info.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
