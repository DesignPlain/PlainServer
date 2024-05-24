package kms

type ReplicaExternalKey struct {
	// The key policy to attach to the KMS key. If you do not specify a key policy, AWS KMS attaches the [default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default) to the KMS key.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// A map of tags to assign to the replica key. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A description of the KMS key.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Specifies whether the replica key is enabled. Disabled KMS keys cannot be used in cryptographic operations. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Base64 encoded 256-bit symmetric encryption key material to import. The KMS key is permanently associated with this key material. The same key material can be [reimported](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html#reimport-key-material), but you cannot import different key material.
	KeyMaterialBase64 string `json:"keyMaterialBase64,omitempty" yaml:"keyMaterialBase64,omitempty"`

	// Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the key becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	ValidTo string `json:"validTo,omitempty" yaml:"validTo,omitempty"`

	/*
	   A flag to indicate whether to bypass the key policy lockout safety check.
	   Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	   For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
	   The default value is `false`.
	*/
	BypassPolicyLockoutSafetyCheck bool `json:"bypassPolicyLockoutSafetyCheck,omitempty" yaml:"bypassPolicyLockoutSafetyCheck,omitempty"`

	/*
	   The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	   If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
	*/
	DeletionWindowInDays int `json:"deletionWindowInDays,omitempty" yaml:"deletionWindowInDays,omitempty"`

	// The ARN of the multi-Region primary key to replicate. The primary key must be in a different AWS Region of the same AWS Partition. You can create only one replica of a given primary key in each AWS Region.
	PrimaryKeyArn string `json:"primaryKeyArn,omitempty" yaml:"primaryKeyArn,omitempty"`
}
