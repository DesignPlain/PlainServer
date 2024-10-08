package kms

type Key struct {
	// Indicates whether the KMS key is a multi-Region (`true`) or regional (`false`) key. Defaults to `false`.
	MultiRegion bool `json:"multiRegion,omitempty" yaml:"multiRegion,omitempty"`

	// ID of the KMS [Custom Key Store](https://docs.aws.amazon.com/kms/latest/developerguide/create-cmk-keystore.html) where the key will be stored instead of KMS (eg CloudHSM).
	CustomKeyStoreId string `json:"customKeyStoreId,omitempty" yaml:"customKeyStoreId,omitempty"`

	/*
	   The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key.
	   If you specify a value, it must be between `7` and `30`, inclusive. If you do not specify a value, it defaults to `30`.
	   If the KMS key is a multi-Region primary key with replicas, the waiting period begins when the last of its replica keys is deleted. Otherwise, the waiting period begins immediately.
	*/
	DeletionWindowInDays int `json:"deletionWindowInDays,omitempty" yaml:"deletionWindowInDays,omitempty"`

	// The description of the key as viewed in AWS console.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Specifies whether [key rotation](http://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html) is enabled. Defaults to `false`.
	EnableKeyRotation bool `json:"enableKeyRotation,omitempty" yaml:"enableKeyRotation,omitempty"`

	// Specifies whether the key is enabled. Defaults to `true`.
	IsEnabled bool `json:"isEnabled,omitempty" yaml:"isEnabled,omitempty"`

	/*
	   Specifies the intended use of the key. Valid values: `ENCRYPT_DECRYPT`, `SIGN_VERIFY`, or `GENERATE_VERIFY_MAC`.
	   Defaults to `ENCRYPT_DECRYPT`.
	*/
	KeyUsage string `json:"keyUsage,omitempty" yaml:"keyUsage,omitempty"`

	/*
	   A valid policy JSON document. Although this is a key policy, not an IAM policy, an `aws.iam.getPolicyDocument`, in the form that designates a principal, can be used.

	   > --NOTE:-- Note: All KMS keys must have a key policy. If a key policy is not specified, AWS gives the KMS key a [default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default) that gives all principals in the owning account unlimited access to all KMS operations for the key. This default key policy effectively delegates all access control to IAM policies and KMS grants.
	*/
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// Custom period of time between each rotation date. Must be a number between 90 and 2560 (inclusive).
	RotationPeriodInDays int `json:"rotationPeriodInDays,omitempty" yaml:"rotationPeriodInDays,omitempty"`

	/*
	   A flag to indicate whether to bypass the key policy lockout safety check.
	   Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.
	   For more information, refer to the scenario in the [Default Key Policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the _AWS Key Management Service Developer Guide_.
	   The default value is `false`.
	*/
	BypassPolicyLockoutSafetyCheck bool `json:"bypassPolicyLockoutSafetyCheck,omitempty" yaml:"bypassPolicyLockoutSafetyCheck,omitempty"`

	/*
	   Specifies whether the key contains a symmetric key or an asymmetric key pair and the encryption algorithms or signing algorithms that the key supports.
	   Valid values: `SYMMETRIC_DEFAULT`,  `RSA_2048`, `RSA_3072`, `RSA_4096`, `HMAC_256`, `ECC_NIST_P256`, `ECC_NIST_P384`, `ECC_NIST_P521`, or `ECC_SECG_P256K1`. Defaults to `SYMMETRIC_DEFAULT`. For help with choosing a key spec, see the [AWS KMS Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html).
	*/
	CustomerMasterKeySpec string `json:"customerMasterKeySpec,omitempty" yaml:"customerMasterKeySpec,omitempty"`

	// A map of tags to assign to the object. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Identifies the external key that serves as key material for the KMS key in an external key store.
	XksKeyId string `json:"xksKeyId,omitempty" yaml:"xksKeyId,omitempty"`
}
