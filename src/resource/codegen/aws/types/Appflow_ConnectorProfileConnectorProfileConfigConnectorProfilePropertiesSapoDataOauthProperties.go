package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOauthProperties struct {
	// The authorization code url required to redirect to SAP Login Page to fetch authorization code for OAuth type authentication.
	AuthCodeUrl string `json:"authCodeUrl,omitempty" yaml:"authCodeUrl,omitempty"`

	// The OAuth scopes required for OAuth type authentication.
	OauthScopes []string `json:"oauthScopes,omitempty" yaml:"oauthScopes,omitempty"`

	//
	TokenUrl string `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`
}
