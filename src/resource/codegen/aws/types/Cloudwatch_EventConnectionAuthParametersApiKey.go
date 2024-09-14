package types

type Cloudwatch_EventConnectionAuthParametersApiKey struct {
	// Header Name.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Header Value. Created and stored in AWS Secrets Manager.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
