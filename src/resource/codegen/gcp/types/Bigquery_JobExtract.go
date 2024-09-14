package types

type Bigquery_JobExtract struct {
	/*
	   The exported file format. Possible values include CSV, NEWLINE_DELIMITED_JSON and AVRO for tables and SAVED_MODEL for models.
	   The default value for tables is CSV. Tables with nested or repeated fields cannot be exported as CSV.
	   The default value for models is SAVED_MODEL.
	*/
	DestinationFormat string `json:"destinationFormat,omitempty" yaml:"destinationFormat,omitempty"`

	// A list of fully-qualified Google Cloud Storage URIs where the extracted table should be written.
	DestinationUris []string `json:"destinationUris,omitempty" yaml:"destinationUris,omitempty"`

	/*
	   When extracting data in CSV format, this defines the delimiter to use between fields in the exported data.
	   Default is ','
	*/
	FieldDelimiter string `json:"fieldDelimiter,omitempty" yaml:"fieldDelimiter,omitempty"`

	// Whether to print out a header row in the results. Default is true.
	PrintHeader bool `json:"printHeader,omitempty" yaml:"printHeader,omitempty"`

	/*
	   A reference to the model being exported.
	   Structure is documented below.
	*/
	SourceModel Bigquery_JobExtractSourceModel `json:"sourceModel,omitempty" yaml:"sourceModel,omitempty"`

	/*
	   A reference to the table being exported.
	   Structure is documented below.
	*/
	SourceTable Bigquery_JobExtractSourceTable `json:"sourceTable,omitempty" yaml:"sourceTable,omitempty"`

	// Whether to use logical types when extracting to AVRO format.
	UseAvroLogicalTypes bool `json:"useAvroLogicalTypes,omitempty" yaml:"useAvroLogicalTypes,omitempty"`

	/*
	   The compression type to use for exported files. Possible values include GZIP, DEFLATE, SNAPPY, and NONE.
	   The default value is NONE. DEFLATE and SNAPPY are only supported for Avro.
	*/
	Compression string `json:"compression,omitempty" yaml:"compression,omitempty"`
}
