package types

type Ec2clientvpn_EndpointConnectionLogOptions struct {
	// The name of the CloudWatch Logs log group.
	CloudwatchLogGroup string `json:"cloudwatchLogGroup,omitempty" yaml:"cloudwatchLogGroup,omitempty"`

	// The name of the CloudWatch Logs log stream to which the connection data is published.
	CloudwatchLogStream string `json:"cloudwatchLogStream,omitempty" yaml:"cloudwatchLogStream,omitempty"`

	// Indicates whether connection logging is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
