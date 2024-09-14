package types

type Glue_getCatalogTableStorageDescriptor struct {
	// Whether the data in the table is compressed.
	Compressed bool `json:"compressed,omitempty" yaml:"compressed,omitempty"`

	// Configuration block for the sort order of each bucket in the table. See `sort_columns` below.
	SortColumns []Glue_getCatalogTableStorageDescriptorSortColumn `json:"sortColumns,omitempty" yaml:"sortColumns,omitempty"`

	// List of reducer grouping columns, clustering columns, and bucketing columns in the table.
	BucketColumns []string `json:"bucketColumns,omitempty" yaml:"bucketColumns,omitempty"`

	// Output format: SequenceFileOutputFormat (binary), or IgnoreKeyTextOutputFormat, or a custom format.
	OutputFormat string `json:"outputFormat,omitempty" yaml:"outputFormat,omitempty"`

	// Configuration block with information about values that appear very frequently in a column (skewed values). See `skewed_info` below.
	SkewedInfos []Glue_getCatalogTableStorageDescriptorSkewedInfo `json:"skewedInfos,omitempty" yaml:"skewedInfos,omitempty"`

	// Configuration block for columns in the table. See `columns` below.
	Columns []Glue_getCatalogTableStorageDescriptorColumn `json:"columns,omitempty" yaml:"columns,omitempty"`

	// Physical location of the table. By default, this takes the form of the warehouse location, followed by the database location in the warehouse, followed by the table name.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Map of initialization parameters for the SerDe, in key-value form.
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Configuration block for serialization and deserialization ("SerDe") information. See `ser_de_info` below.
	SerDeInfos []Glue_getCatalogTableStorageDescriptorSerDeInfo `json:"serDeInfos,omitempty" yaml:"serDeInfos,omitempty"`

	// List of locations that point to the path where a Delta table is located
	AdditionalLocations []string `json:"additionalLocations,omitempty" yaml:"additionalLocations,omitempty"`

	// Input format: SequenceFileInputFormat (binary), or TextInputFormat, or a custom format.
	InputFormat string `json:"inputFormat,omitempty" yaml:"inputFormat,omitempty"`

	// Is if the table contains any dimension columns.
	NumberOfBuckets int `json:"numberOfBuckets,omitempty" yaml:"numberOfBuckets,omitempty"`

	// Object that references a schema stored in the AWS Glue Schema Registry. See `schema_reference` below.
	SchemaReferences []Glue_getCatalogTableStorageDescriptorSchemaReference `json:"schemaReferences,omitempty" yaml:"schemaReferences,omitempty"`

	// Whether the table data is stored in subdirectories.
	StoredAsSubDirectories bool `json:"storedAsSubDirectories,omitempty" yaml:"storedAsSubDirectories,omitempty"`
}
