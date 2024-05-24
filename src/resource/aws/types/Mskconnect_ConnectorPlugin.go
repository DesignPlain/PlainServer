package types

type Mskconnect_ConnectorPlugin struct {
	// Details about a custom plugin. See below.
	CustomPlugin Mskconnect_ConnectorPluginCustomPlugin `json:"customPlugin,omitempty" yaml:"customPlugin,omitempty"`
}
