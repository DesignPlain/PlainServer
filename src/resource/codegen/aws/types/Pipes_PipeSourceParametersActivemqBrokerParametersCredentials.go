package types

type Pipes_PipeSourceParametersActivemqBrokerParametersCredentials struct {
	// The ARN of the Secrets Manager secret containing the credentials.
	BasicAuth string `json:"basicAuth,omitempty" yaml:"basicAuth,omitempty"`
}
