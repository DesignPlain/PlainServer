package essentialcontacts

import types "libds/gcp/types"

type DocumentAiWarehouseDocumentSchema struct {
	// Name of the schema given by the user.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Tells whether the document is a folder or a typical document.
	DocumentIsFolder bool `json:"documentIsFolder,omitempty" yaml:"documentIsFolder,omitempty"`

	// The location of the resource.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The unique identifier of the project.
	ProjectNumber string `json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`

	/*
	   Defines the metadata for a schema property.
	   Structure is documented below.
	*/
	PropertyDefinitions []types.Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinition `json:"propertyDefinitions,omitempty" yaml:"propertyDefinitions,omitempty"`
}
