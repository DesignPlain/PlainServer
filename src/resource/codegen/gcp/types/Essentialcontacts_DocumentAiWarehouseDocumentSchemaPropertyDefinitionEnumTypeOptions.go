package types

type Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinitionEnumTypeOptions struct {
	// List of possible enum values.
	PossibleValues []string `json:"possibleValues,omitempty" yaml:"possibleValues,omitempty"`

	/*
	   Make sure the enum property value provided in the document is in the possile value list during document creation. The validation check runs by default.

	   - - -
	*/
	ValidationCheckDisabled bool `json:"validationCheckDisabled,omitempty" yaml:"validationCheckDisabled,omitempty"`
}
