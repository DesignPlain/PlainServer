package verifiedpermissions

import types "libds/aws/types"

type PolicyStore struct {
	// A description of the Policy Store.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Validation settings for the policy store.
	ValidationSettings types.Verifiedpermissions_PolicyStoreValidationSettings `json:"validationSettings,omitempty" yaml:"validationSettings,omitempty"`
}
