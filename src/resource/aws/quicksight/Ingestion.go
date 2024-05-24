package quicksight

type Ingestion struct {
	// AWS account ID.
	AwsAccountId string `json:"awsAccountId,omitempty" yaml:"awsAccountId,omitempty"`

	// ID of the dataset used in the ingestion.
	DataSetId string `json:"dataSetId,omitempty" yaml:"dataSetId,omitempty"`

	// ID for the ingestion.
	IngestionId string `json:"ingestionId,omitempty" yaml:"ingestionId,omitempty"`

	/*
	   Type of ingestion to be created. Valid values are `INCREMENTAL_REFRESH` and `FULL_REFRESH`.

	   The following arguments are optional:
	*/
	IngestionType string `json:"ingestionType,omitempty" yaml:"ingestionType,omitempty"`
}
