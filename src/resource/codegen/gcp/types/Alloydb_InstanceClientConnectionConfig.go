package types

type Alloydb_InstanceClientConnectionConfig struct {
	// Configuration to enforce connectors only (ex: AuthProxy) connections to the database.
	RequireConnectors bool `json:"requireConnectors,omitempty" yaml:"requireConnectors,omitempty"`

	/*
	   SSL config option for this instance.
	   Structure is documented below.
	*/
	SslConfig Alloydb_InstanceClientConnectionConfigSslConfig `json:"sslConfig,omitempty" yaml:"sslConfig,omitempty"`
}
