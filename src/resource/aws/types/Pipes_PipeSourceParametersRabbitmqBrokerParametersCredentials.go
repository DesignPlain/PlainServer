package types

type Pipes_PipeSourceParametersRabbitmqBrokerParametersCredentials struct {
	// The ARN of the Secrets Manager secret containing the credentials.
	BasicAuth string `json:"basicAuth,omitempty" yaml:"basicAuth,omitempty"`
}
