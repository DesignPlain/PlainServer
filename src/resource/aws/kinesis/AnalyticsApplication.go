package kinesis

import types "DesignSphere_Server/src/resource/aws/types"

type AnalyticsApplication struct {
	// SQL Code to transform input data, and generate output.
	Code string `json:"code,omitempty" yaml:"code,omitempty"`

	// Output destination configuration of the application. See Outputs below for more details.
	Outputs []types.Kinesis_AnalyticsApplicationOutput `json:"outputs,omitempty" yaml:"outputs,omitempty"`

	/*
	   Whether to start or stop the Kinesis Analytics Application. To start an application, an input with a defined `starting_position` must be configured.
	   To modify an application's starting position, first stop the application by setting `start_application = false`, then update `starting_position` and set `start_application = true`.
	*/
	StartApplication bool `json:"startApplication,omitempty" yaml:"startApplication,omitempty"`

	/*
	   The CloudWatch log stream options to monitor application errors.
	   See CloudWatch Logging Options below for more details.
	*/
	CloudwatchLoggingOptions types.Kinesis_AnalyticsApplicationCloudwatchLoggingOptions `json:"cloudwatchLoggingOptions,omitempty" yaml:"cloudwatchLoggingOptions,omitempty"`

	// Description of the application.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Input configuration of the application. See Inputs below for more details.
	Inputs types.Kinesis_AnalyticsApplicationInputs `json:"inputs,omitempty" yaml:"inputs,omitempty"`

	// Name of the Kinesis Analytics Application.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   An S3 Reference Data Source for the application.
	   See Reference Data Sources below for more details.
	*/
	ReferenceDataSources types.Kinesis_AnalyticsApplicationReferenceDataSources `json:"referenceDataSources,omitempty" yaml:"referenceDataSources,omitempty"`

	// Key-value map of tags for the Kinesis Analytics Application. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
