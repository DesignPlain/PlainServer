package verifiedpermissions

import types "DesignSphere_Server/src/resource/aws/types"

type PolicyStore struct {
	// Validation settings for the policy store.
	ValidationSettings types.Verifiedpermissions_PolicyStoreValidationSettings `json:"validationSettings,omitempty" yaml:"validationSettings,omitempty"`

	// A description of the Policy Store.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
