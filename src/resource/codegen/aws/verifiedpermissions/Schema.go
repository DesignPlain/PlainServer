package verifiedpermissions

import types "libds/aws/types"

type Schema struct {
	// The definition of the schema.
	Definition types.Verifiedpermissions_SchemaDefinition `json:"definition,omitempty" yaml:"definition,omitempty"`

	// The ID of the Policy Store.
	PolicyStoreId string `json:"policyStoreId,omitempty" yaml:"policyStoreId,omitempty"`
}
