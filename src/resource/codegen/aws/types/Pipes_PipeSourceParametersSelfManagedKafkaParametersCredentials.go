package types

type Pipes_PipeSourceParametersSelfManagedKafkaParametersCredentials struct {
	// The ARN of the Secrets Manager secret containing the credentials.
	BasicAuth string `json:"basicAuth,omitempty" yaml:"basicAuth,omitempty"`

	// The ARN of the Secrets Manager secret containing the credentials.
	ClientCertificateTlsAuth string `json:"clientCertificateTlsAuth,omitempty" yaml:"clientCertificateTlsAuth,omitempty"`

	// The ARN of the Secrets Manager secret containing the credentials.
	SaslScram256Auth string `json:"saslScram256Auth,omitempty" yaml:"saslScram256Auth,omitempty"`

	// The ARN of the Secrets Manager secret containing the credentials.
	SaslScram512Auth string `json:"saslScram512Auth,omitempty" yaml:"saslScram512Auth,omitempty"`
}
