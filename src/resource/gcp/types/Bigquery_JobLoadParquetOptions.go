package types

type Bigquery_JobLoadParquetOptions struct {
	// If sourceFormat is set to PARQUET, indicates whether to use schema inference specifically for Parquet LIST logical type.
	EnableListInference bool `json:"enableListInference,omitempty" yaml:"enableListInference,omitempty"`

	// If sourceFormat is set to PARQUET, indicates whether to infer Parquet ENUM logical type as STRING instead of BYTES by default.
	EnumAsString bool `json:"enumAsString,omitempty" yaml:"enumAsString,omitempty"`
}
