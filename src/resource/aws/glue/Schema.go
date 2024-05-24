package glue

type Schema struct {
	// The data format of the schema definition. Valid values are `AVRO`, `JSON` and `PROTOBUF`.
	DataFormat string `json:"dataFormat,omitempty" yaml:"dataFormat,omitempty"`

	// A description of the schema.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The ARN of the Glue Registry to create the schema in.
	RegistryArn string `json:"registryArn,omitempty" yaml:"registryArn,omitempty"`

	// The schema definition using the `data_format` setting for `schema_name`.
	SchemaDefinition string `json:"schemaDefinition,omitempty" yaml:"schemaDefinition,omitempty"`

	// The Name of the schema.
	SchemaName string `json:"schemaName,omitempty" yaml:"schemaName,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The compatibility mode of the schema. Values values are: `NONE`, `DISABLED`, `BACKWARD`, `BACKWARD_ALL`, `FORWARD`, `FORWARD_ALL`, `FULL`, and `FULL_ALL`.
	Compatibility string `json:"compatibility,omitempty" yaml:"compatibility,omitempty"`
}
