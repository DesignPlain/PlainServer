package types

type Sfn_ActivityEncryptionConfiguration struct {
	// Maximum duration for which Activities will reuse data keys. When the period expires, Activities will call GenerateDataKey. This setting only applies to customer managed KMS key and does not apply to AWS owned KMS key.
	KmsDataKeyReusePeriodSeconds int `json:"kmsDataKeyReusePeriodSeconds,omitempty" yaml:"kmsDataKeyReusePeriodSeconds,omitempty"`

	// The alias, alias ARN, key ID, or key ARN of the symmetric encryption KMS key that encrypts the data key. To specify a KMS key in a different AWS account, the customer must use the key ARN or alias ARN. For more information regarding kms_key_id, see [KeyId](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters) in the KMS documentation.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// The encryption option specified for the activity. Valid values: `AWS_KMS_KEY`, `CUSTOMER_MANAGED_KMS_KEY`
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
