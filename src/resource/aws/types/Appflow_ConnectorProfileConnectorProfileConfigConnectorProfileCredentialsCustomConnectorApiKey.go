package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorApiKey struct {
	// Unique alphanumeric identifier used to authenticate a user, developer, or calling program to your API.
	ApiKey string `json:"apiKey,omitempty" yaml:"apiKey,omitempty"`

	// The API secret key required for API key authentication.
	ApiSecretKey string `json:"apiSecretKey,omitempty" yaml:"apiSecretKey,omitempty"`
}
