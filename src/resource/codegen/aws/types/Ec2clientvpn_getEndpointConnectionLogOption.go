package types

type Ec2clientvpn_getEndpointConnectionLogOption struct {
	//
	CloudwatchLogGroup string `json:"cloudwatchLogGroup,omitempty" yaml:"cloudwatchLogGroup,omitempty"`

	//
	CloudwatchLogStream string `json:"cloudwatchLogStream,omitempty" yaml:"cloudwatchLogStream,omitempty"`

	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
