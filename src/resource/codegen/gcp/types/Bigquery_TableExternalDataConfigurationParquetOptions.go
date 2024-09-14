package types

type Bigquery_TableExternalDataConfigurationParquetOptions struct {
	// Indicates whether to use schema inference specifically for Parquet LIST logical type.
	EnableListInference bool `json:"enableListInference,omitempty" yaml:"enableListInference,omitempty"`

	// Indicates whether to infer Parquet ENUM logical type as STRING instead of BYTES by default.
	EnumAsString bool `json:"enumAsString,omitempty" yaml:"enumAsString,omitempty"`
}
