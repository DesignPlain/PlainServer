package types

type Glue_PartitionStorageDescriptor struct {
	// A list of reducer grouping columns, clustering columns, and bucketing columns in the table.
	BucketColumns []string `json:"bucketColumns,omitempty" yaml:"bucketColumns,omitempty"`

	// The input format: SequenceFileInputFormat (binary), or TextInputFormat, or a custom format.
	InputFormat string `json:"inputFormat,omitempty" yaml:"inputFormat,omitempty"`

	// The physical location of the table. By default this takes the form of the warehouse location, followed by the database location in the warehouse, followed by the table name.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// User-supplied properties in key-value form.
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Serialization/deserialization (SerDe) information.
	SerDeInfo Glue_PartitionStorageDescriptorSerDeInfo `json:"serDeInfo,omitempty" yaml:"serDeInfo,omitempty"`

	// Information about values that appear very frequently in a column (skewed values).
	SkewedInfo Glue_PartitionStorageDescriptorSkewedInfo `json:"skewedInfo,omitempty" yaml:"skewedInfo,omitempty"`

	// A list of Order objects specifying the sort order of each bucket in the table.
	SortColumns []Glue_PartitionStorageDescriptorSortColumn `json:"sortColumns,omitempty" yaml:"sortColumns,omitempty"`

	// True if the table data is stored in subdirectories, or False if not.
	StoredAsSubDirectories bool `json:"storedAsSubDirectories,omitempty" yaml:"storedAsSubDirectories,omitempty"`

	// A list of the Columns in the table.
	Columns []Glue_PartitionStorageDescriptorColumn `json:"columns,omitempty" yaml:"columns,omitempty"`

	// True if the data in the table is compressed, or False if not.
	Compressed bool `json:"compressed,omitempty" yaml:"compressed,omitempty"`

	// Must be specified if the table contains any dimension columns.
	NumberOfBuckets int `json:"numberOfBuckets,omitempty" yaml:"numberOfBuckets,omitempty"`

	// The output format: SequenceFileOutputFormat (binary), or IgnoreKeyTextOutputFormat, or a custom format.
	OutputFormat string `json:"outputFormat,omitempty" yaml:"outputFormat,omitempty"`
}
