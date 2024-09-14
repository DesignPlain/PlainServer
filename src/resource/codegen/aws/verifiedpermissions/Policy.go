package verifiedpermissions

import types "libds/aws/types"

type Policy struct {
	// The definition of the policy. See Definition below.
	Definition types.Verifiedpermissions_PolicyDefinition `json:"definition,omitempty" yaml:"definition,omitempty"`

	// The Policy Store ID of the policy store.
	PolicyStoreId string `json:"policyStoreId,omitempty" yaml:"policyStoreId,omitempty"`
}
