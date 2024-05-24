package types

type Cloudwatch_EventConnectionAuthParametersInvocationHttpParametersHeader struct {
	// Specified whether the value is secret.
	IsValueSecret bool `json:"isValueSecret,omitempty" yaml:"isValueSecret,omitempty"`

	// Header Name.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Header Value. Created and stored in AWS Secrets Manager.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
