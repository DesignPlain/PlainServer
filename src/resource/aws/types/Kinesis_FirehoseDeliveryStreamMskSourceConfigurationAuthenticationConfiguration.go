package types

type Kinesis_FirehoseDeliveryStreamMskSourceConfigurationAuthenticationConfiguration struct {
	// The type of connectivity used to access the Amazon MSK cluster. Valid values: `PUBLIC`, `PRIVATE`.
	Connectivity string `json:"connectivity,omitempty" yaml:"connectivity,omitempty"`

	// The ARN of the role used to access the Amazon MSK cluster.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
