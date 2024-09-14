package types

type Glue_getCatalogTableStorageDescriptorSchemaReference struct {
	// Configuration block that contains schema identity fields. See `schema_id` below.
	SchemaIds []Glue_getCatalogTableStorageDescriptorSchemaReferenceSchemaId `json:"schemaIds,omitempty" yaml:"schemaIds,omitempty"`

	// Unique ID assigned to a version of the schema.
	SchemaVersionId string `json:"schemaVersionId,omitempty" yaml:"schemaVersionId,omitempty"`

	// Version number of the schema.
	SchemaVersionNumber int `json:"schemaVersionNumber,omitempty" yaml:"schemaVersionNumber,omitempty"`
}
