package types

type Sagemaker_DataQualityJobDefinitionDataQualityJobInputBatchTransformInputDatasetFormat struct {
	// The CSV dataset used in the monitoring job. Fields are documented below.
	Csv Sagemaker_DataQualityJobDefinitionDataQualityJobInputBatchTransformInputDatasetFormatCsv `json:"csv,omitempty" yaml:"csv,omitempty"`

	// The JSON dataset used in the monitoring job. Fields are documented below.
	Json Sagemaker_DataQualityJobDefinitionDataQualityJobInputBatchTransformInputDatasetFormatJson `json:"json,omitempty" yaml:"json,omitempty"`
}
