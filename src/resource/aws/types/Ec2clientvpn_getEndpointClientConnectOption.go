package types

type Ec2clientvpn_getEndpointClientConnectOption struct {
	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	//
	LambdaFunctionArn string `json:"lambdaFunctionArn,omitempty" yaml:"lambdaFunctionArn,omitempty"`
}
