package types

type Bigquery_TableExternalDataConfigurationAvroOptions struct {
	/*
	   If is set to true, indicates whether
	   to interpret logical types as the corresponding BigQuery data type
	   (for example, TIMESTAMP), instead of using the raw type (for example, INTEGER).
	*/
	UseAvroLogicalTypes bool `json:"useAvroLogicalTypes,omitempty" yaml:"useAvroLogicalTypes,omitempty"`
}
