package types

type Kinesis_AnalyticsApplicationInputsProcessingConfiguration struct {
	// The Lambda function configuration. See Lambda below for more details.
	Lambda Kinesis_AnalyticsApplicationInputsProcessingConfigurationLambda `json:"lambda,omitempty" yaml:"lambda,omitempty"`
}
