package types

type Cloudwatch_EventConnectionAuthParametersInvocationHttpParametersBody struct {
	// Specified whether the value is secret.
	IsValueSecret bool `json:"isValueSecret,omitempty" yaml:"isValueSecret,omitempty"`

	// The key for the parameter.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// The value associated with the key. Created and stored in AWS Secrets Manager if is secret.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
