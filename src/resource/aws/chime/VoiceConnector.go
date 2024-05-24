package chime

type VoiceConnector struct {
	// The AWS Region in which the Amazon Chime Voice Connector is created. Default value: `us-east-1`
	AwsRegion string `json:"awsRegion,omitempty" yaml:"awsRegion,omitempty"`

	// The name of the Amazon Chime Voice Connector.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   When enabled, requires encryption for the Amazon Chime Voice Connector.

	   The following arguments are optional:
	*/
	RequireEncryption bool `json:"requireEncryption,omitempty" yaml:"requireEncryption,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
