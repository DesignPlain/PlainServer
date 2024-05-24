package types

type Glue_CatalogTableStorageDescriptorSchemaReference struct {
	// Configuration block that contains schema identity fields. Either this or the `schema_version_id` has to be provided. See `schema_id` below.
	SchemaId Glue_CatalogTableStorageDescriptorSchemaReferenceSchemaId `json:"schemaId,omitempty" yaml:"schemaId,omitempty"`

	// Unique ID assigned to a version of the schema. Either this or the `schema_id` has to be provided.
	SchemaVersionId string `json:"schemaVersionId,omitempty" yaml:"schemaVersionId,omitempty"`

	// Version number of the schema.
	SchemaVersionNumber int `json:"schemaVersionNumber,omitempty" yaml:"schemaVersionNumber,omitempty"`
}
