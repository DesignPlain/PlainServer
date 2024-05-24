package types

type Glue_CatalogTableStorageDescriptorSchemaReferenceSchemaId struct {
	// Name of the schema registry that contains the schema. Must be provided when `schema_name` is specified and conflicts with `schema_arn`.
	RegistryName string `json:"registryName,omitempty" yaml:"registryName,omitempty"`

	// ARN of the schema. One of `schema_arn` or `schema_name` has to be provided.
	SchemaArn string `json:"schemaArn,omitempty" yaml:"schemaArn,omitempty"`

	// Name of the schema. One of `schema_arn` or `schema_name` has to be provided.
	SchemaName string `json:"schemaName,omitempty" yaml:"schemaName,omitempty"`
}
