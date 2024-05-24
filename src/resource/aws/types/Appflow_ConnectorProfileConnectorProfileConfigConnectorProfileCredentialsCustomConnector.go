package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnector struct {
	// If the connector uses the custom authentication mechanism, this holds the required credentials.
	Custom Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorCustom `json:"custom,omitempty" yaml:"custom,omitempty"`

	// OAuth 2.0 credentials required for the authentication of the user.
	Oauth2 Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2 `json:"oauth2,omitempty" yaml:"oauth2,omitempty"`

	// Unique alphanumeric identifier used to authenticate a user, developer, or calling program to your API.
	ApiKey Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorApiKey `json:"apiKey,omitempty" yaml:"apiKey,omitempty"`

	// The authentication type that the custom connector uses for authenticating while creating a connector profile. One of: `APIKEY`, `BASIC`, `CUSTOM`, `OAUTH2`.
	AuthenticationType string `json:"authenticationType,omitempty" yaml:"authenticationType,omitempty"`

	// Basic credentials that are required for the authentication of the user.
	Basic Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorBasic `json:"basic,omitempty" yaml:"basic,omitempty"`
}
