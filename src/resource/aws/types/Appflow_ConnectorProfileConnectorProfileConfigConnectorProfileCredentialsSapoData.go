package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoData struct {
	// The SAPOData basic authentication credentials.
	BasicAuthCredentials Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoDataBasicAuthCredentials `json:"basicAuthCredentials,omitempty" yaml:"basicAuthCredentials,omitempty"`

	// The SAPOData OAuth type authentication credentials.
	OauthCredentials Appflow_ConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoDataOauthCredentials `json:"oauthCredentials,omitempty" yaml:"oauthCredentials,omitempty"`
}
