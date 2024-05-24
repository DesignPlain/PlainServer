package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoData struct {
	// The application path to catalog service.
	ApplicationServicePath string `json:"applicationServicePath,omitempty" yaml:"applicationServicePath,omitempty"`

	// The client number for the client creating the connection.
	ClientNumber string `json:"clientNumber,omitempty" yaml:"clientNumber,omitempty"`

	// The logon language of SAPOData instance.
	LogonLanguage string `json:"logonLanguage,omitempty" yaml:"logonLanguage,omitempty"`

	// The SAPOData OAuth properties required for OAuth type authentication.
	OauthProperties Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOauthProperties `json:"oauthProperties,omitempty" yaml:"oauthProperties,omitempty"`

	// The port number of the SAPOData instance.
	PortNumber int `json:"portNumber,omitempty" yaml:"portNumber,omitempty"`

	// The SAPOData Private Link service name to be used for private data transfers.
	PrivateLinkServiceName string `json:"privateLinkServiceName,omitempty" yaml:"privateLinkServiceName,omitempty"`

	// The location of the SAPOData resource.
	ApplicationHostUrl string `json:"applicationHostUrl,omitempty" yaml:"applicationHostUrl,omitempty"`
}
