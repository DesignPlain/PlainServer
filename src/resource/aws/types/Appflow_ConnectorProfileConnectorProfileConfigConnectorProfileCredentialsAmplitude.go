package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitude struct {
	// Unique alphanumeric identifier used to authenticate a user, developer, or calling program to your API.
	ApiKey string `json:"apiKey,omitempty" yaml:"apiKey,omitempty"`

	// The Secret Access Key portion of the credentials.
	SecretKey string `json:"secretKey,omitempty" yaml:"secretKey,omitempty"`
}
