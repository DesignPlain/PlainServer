package verifiedpermissions

import types "libds/aws/types"

type IdentitySource struct {
	// Specifies the details required to communicate with the identity provider (IdP) associated with this identity source. See Configuration below.
	Configuration types.Verifiedpermissions_IdentitySourceConfiguration `json:"configuration,omitempty" yaml:"configuration,omitempty"`

	// Specifies the ID of the policy store in which you want to store this identity source.
	PolicyStoreId string `json:"policyStoreId,omitempty" yaml:"policyStoreId,omitempty"`

	// Specifies the namespace and data type of the principals generated for identities authenticated by the new identity source.
	PrincipalEntityType string `json:"principalEntityType,omitempty" yaml:"principalEntityType,omitempty"`
}
