package types

type Ec2clientvpn_EndpointClientConnectOptions struct {
	// Indicates whether client connect options are enabled. The default is `false` (not enabled).
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The Amazon Resource Name (ARN) of the Lambda function used for connection authorization.
	LambdaFunctionArn string `json:"lambdaFunctionArn,omitempty" yaml:"lambdaFunctionArn,omitempty"`
}
