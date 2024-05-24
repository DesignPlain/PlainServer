package types

type Sagemaker_DataQualityJobDefinitionDataQualityAppSpecification struct {
	// Sets the environment variables in the container that the monitoring job runs. A list of key value pairs.
	Environment map[string]string `json:"environment,omitempty" yaml:"environment,omitempty"`

	// The container image that the data quality monitoring job runs.
	ImageUri string `json:"imageUri,omitempty" yaml:"imageUri,omitempty"`

	// An Amazon S3 URI to a script that is called after analysis has been performed. Applicable only for the built-in (first party) containers.
	PostAnalyticsProcessorSourceUri string `json:"postAnalyticsProcessorSourceUri,omitempty" yaml:"postAnalyticsProcessorSourceUri,omitempty"`

	// An Amazon S3 URI to a script that is called per row prior to running analysis. It can base64 decode the payload and convert it into a flatted json so that the built-in container can use the converted data. Applicable only for the built-in (first party) containers.
	RecordPreprocessorSourceUri string `json:"recordPreprocessorSourceUri,omitempty" yaml:"recordPreprocessorSourceUri,omitempty"`
}
