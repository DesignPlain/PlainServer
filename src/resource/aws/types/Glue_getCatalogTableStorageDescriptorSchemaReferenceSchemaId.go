package types

type Glue_getCatalogTableStorageDescriptorSchemaReferenceSchemaId struct {
	// Name of the schema registry that contains the schema.
	RegistryName string `json:"registryName,omitempty" yaml:"registryName,omitempty"`

	// ARN of the schema.
	SchemaArn string `json:"schemaArn,omitempty" yaml:"schemaArn,omitempty"`

	// Name of the schema.
	SchemaName string `json:"schemaName,omitempty" yaml:"schemaName,omitempty"`
}
