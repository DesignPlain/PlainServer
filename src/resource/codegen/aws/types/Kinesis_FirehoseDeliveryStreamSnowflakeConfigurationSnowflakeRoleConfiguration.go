package types

type Kinesis_FirehoseDeliveryStreamSnowflakeConfigurationSnowflakeRoleConfiguration struct {
	// Whether the Snowflake role is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The Snowflake role.
	SnowflakeRole string `json:"snowflakeRole,omitempty" yaml:"snowflakeRole,omitempty"`
}
