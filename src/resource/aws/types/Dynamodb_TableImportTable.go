package types

type Dynamodb_TableImportTable struct {
	/*
	   Type of compression to be used on the input coming from the imported table.
	   Valid values are `GZIP`, `ZSTD` and `NONE`.
	*/
	InputCompressionType string `json:"inputCompressionType,omitempty" yaml:"inputCompressionType,omitempty"`

	/*
	   The format of the source data.
	   Valid values are `CSV`, `DYNAMODB_JSON`, and `ION`.
	*/
	InputFormat string `json:"inputFormat,omitempty" yaml:"inputFormat,omitempty"`

	/*
	   Describe the format options for the data that was imported into the target table.
	   There is one value, `csv`.
	   See below.
	*/
	InputFormatOptions Dynamodb_TableImportTableInputFormatOptions `json:"inputFormatOptions,omitempty" yaml:"inputFormatOptions,omitempty"`

	/*
	   Values for the S3 bucket the source file is imported from.
	   See below.
	*/
	S3BucketSource Dynamodb_TableImportTableS3BucketSource `json:"s3BucketSource,omitempty" yaml:"s3BucketSource,omitempty"`
}
