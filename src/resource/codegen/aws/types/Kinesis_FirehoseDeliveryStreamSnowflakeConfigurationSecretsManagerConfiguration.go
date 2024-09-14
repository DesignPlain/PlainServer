package types

type Kinesis_FirehoseDeliveryStreamSnowflakeConfigurationSecretsManagerConfiguration struct {
	// Enables or disables the Secrets Manager configuration.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The ARN of the role the stream assumes.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The ARN of the Secrets Manager secret. This value is required if `enabled` is true.
	SecretArn string `json:"secretArn,omitempty" yaml:"secretArn,omitempty"`
}
