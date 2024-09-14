package athena

type DataCatalog struct {
	// Description of the data catalog.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of the data catalog. The catalog name must be unique for the AWS account and can use a maximum of 128 alphanumeric, underscore, at sign, or hyphen characters.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key value pairs that specifies the Lambda function or functions to use for the data catalog. The mapping used depends on the catalog type.
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Type of data catalog: `LAMBDA` for a federated catalog, `GLUE` for AWS Glue Catalog, or `HIVE` for an external hive metastore.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
