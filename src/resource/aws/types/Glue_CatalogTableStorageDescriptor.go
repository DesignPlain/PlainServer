package types

type Glue_CatalogTableStorageDescriptor struct {
	// List of reducer grouping columns, clustering columns, and bucketing columns in the table.
	BucketColumns []string `json:"bucketColumns,omitempty" yaml:"bucketColumns,omitempty"`

	// Configuration block for columns in the table. See `columns` below.
	Columns []Glue_CatalogTableStorageDescriptorColumn `json:"columns,omitempty" yaml:"columns,omitempty"`

	// Whether the data in the table is compressed.
	Compressed bool `json:"compressed,omitempty" yaml:"compressed,omitempty"`

	// Input format: SequenceFileInputFormat (binary), or TextInputFormat, or a custom format.
	InputFormat string `json:"inputFormat,omitempty" yaml:"inputFormat,omitempty"`

	// Output format: SequenceFileOutputFormat (binary), or IgnoreKeyTextOutputFormat, or a custom format.
	OutputFormat string `json:"outputFormat,omitempty" yaml:"outputFormat,omitempty"`

	// User-supplied properties in key-value form.
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Configuration block with information about values that appear very frequently in a column (skewed values). See `skewed_info` below.
	SkewedInfo Glue_CatalogTableStorageDescriptorSkewedInfo `json:"skewedInfo,omitempty" yaml:"skewedInfo,omitempty"`

	// Whether the table data is stored in subdirectories.
	StoredAsSubDirectories bool `json:"storedAsSubDirectories,omitempty" yaml:"storedAsSubDirectories,omitempty"`

	// Physical location of the table. By default this takes the form of the warehouse location, followed by the database location in the warehouse, followed by the table name.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Must be specified if the table contains any dimension columns.
	NumberOfBuckets int `json:"numberOfBuckets,omitempty" yaml:"numberOfBuckets,omitempty"`

	// Object that references a schema stored in the AWS Glue Schema Registry. When creating a table, you can pass an empty list of columns for the schema, and instead use a schema reference. See Schema Reference below.
	SchemaReference Glue_CatalogTableStorageDescriptorSchemaReference `json:"schemaReference,omitempty" yaml:"schemaReference,omitempty"`

	// Configuration block for serialization and deserialization ("SerDe") information. See `ser_de_info` below.
	SerDeInfo Glue_CatalogTableStorageDescriptorSerDeInfo `json:"serDeInfo,omitempty" yaml:"serDeInfo,omitempty"`

	// Configuration block for the sort order of each bucket in the table. See `sort_columns` below.
	SortColumns []Glue_CatalogTableStorageDescriptorSortColumn `json:"sortColumns,omitempty" yaml:"sortColumns,omitempty"`
}
