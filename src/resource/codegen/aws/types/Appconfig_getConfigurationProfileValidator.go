package types

type Appconfig_getConfigurationProfileValidator struct {
	// Either the JSON Schema content or the ARN of an AWS Lambda function.
	Content string `json:"content,omitempty" yaml:"content,omitempty"`

	// Type of validator. Valid values: JSON_SCHEMA and LAMBDA.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
