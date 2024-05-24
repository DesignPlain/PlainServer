package types

type Pipes_PipeSourceParametersManagedStreamingKafkaParametersCredentials struct {
	// The ARN of the Secrets Manager secret containing the credentials.
	SaslScram512Auth string `json:"saslScram512Auth,omitempty" yaml:"saslScram512Auth,omitempty"`

	// The ARN of the Secrets Manager secret containing the credentials.
	ClientCertificateTlsAuth string `json:"clientCertificateTlsAuth,omitempty" yaml:"clientCertificateTlsAuth,omitempty"`
}
