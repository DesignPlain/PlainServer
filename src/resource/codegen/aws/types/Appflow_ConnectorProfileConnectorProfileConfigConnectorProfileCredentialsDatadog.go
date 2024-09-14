package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDatadog struct {
	//
	ApiKey string `json:"apiKey,omitempty" yaml:"apiKey,omitempty"`

	// Application keys, in conjunction with your API key, give you full access to Datadog’s programmatic API. Application keys are associated with the user account that created them. The application key is used to log all requests made to the API.
	ApplicationKey string `json:"applicationKey,omitempty" yaml:"applicationKey,omitempty"`
}
