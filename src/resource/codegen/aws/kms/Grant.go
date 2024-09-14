package kms

import types "libds/aws/types"

type Grant struct {
	/*
	   If set to false (the default) the grants will be revoked upon deletion, and if set to true the grants will try to be retired upon deletion. Note that retiring grants requires special permissions, hence why we default to revoking grants.
	   See [RetireGrant](https://docs.aws.amazon.com/kms/latest/APIReference/API_RetireGrant.html) for more information.
	*/
	RetireOnDelete bool `json:"retireOnDelete,omitempty" yaml:"retireOnDelete,omitempty"`

	// The principal that is given permission to retire the grant by using RetireGrant operation in ARN format. Note that due to eventual consistency issues around IAM principals, the providers's state may not always be refreshed to reflect what is true in AWS.
	RetiringPrincipal string `json:"retiringPrincipal,omitempty" yaml:"retiringPrincipal,omitempty"`

	// A structure that you can use to allow certain operations in the grant only when the desired encryption context is present. For more information about encryption context, see [Encryption Context](http://docs.aws.amazon.com/kms/latest/developerguide/encryption-context.html).
	Constraints []types.Kms_GrantConstraint `json:"constraints,omitempty" yaml:"constraints,omitempty"`

	// A list of grant tokens to be used when creating the grant. See [Grant Tokens](http://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token) for more information about grant tokens.
	GrantCreationTokens []string `json:"grantCreationTokens,omitempty" yaml:"grantCreationTokens,omitempty"`

	// The principal that is given permission to perform the operations that the grant permits in ARN format. Note that due to eventual consistency issues around IAM principals, the providers's state may not always be refreshed to reflect what is true in AWS.
	GranteePrincipal string `json:"granteePrincipal,omitempty" yaml:"granteePrincipal,omitempty"`

	// The unique identifier for the customer master key (CMK) that the grant applies to. Specify the key ID or the Amazon Resource Name (ARN) of the CMK. To specify a CMK in a different AWS account, you must use the key ARN.
	KeyId string `json:"keyId,omitempty" yaml:"keyId,omitempty"`

	// A friendly name for identifying the grant.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A list of operations that the grant permits. The permitted values are: `Decrypt`, `Encrypt`, `GenerateDataKey`, `GenerateDataKeyWithoutPlaintext`, `ReEncryptFrom`, `ReEncryptTo`, `Sign`, `Verify`, `GetPublicKey`, `CreateGrant`, `RetireGrant`, `DescribeKey`, `GenerateDataKeyPair`, or `GenerateDataKeyPairWithoutPlaintext`.
	Operations []string `json:"operations,omitempty" yaml:"operations,omitempty"`
}
