package types

type Integrationconnectors_ConnectionSslConfig struct {
	/*
	   Type of Server Cert (PEM/JKS/.. etc.)
	   Possible values are: `PEM`.
	*/
	ServerCertType string `json:"serverCertType,omitempty" yaml:"serverCertType,omitempty"`

	/*
	   Additional SSL related field values.
	   Structure is documented below.
	*/
	AdditionalVariables []Integrationconnectors_ConnectionSslConfigAdditionalVariable `json:"additionalVariables,omitempty" yaml:"additionalVariables,omitempty"`

	/*
	   Client Certificate
	   Structure is documented below.
	*/
	ClientCertificate Integrationconnectors_ConnectionSslConfigClientCertificate `json:"clientCertificate,omitempty" yaml:"clientCertificate,omitempty"`

	/*
	   Private Server Certificate. Needs to be specified if trust model is PRIVATE.
	   Structure is documented below.
	*/
	PrivateServerCertificate Integrationconnectors_ConnectionSslConfigPrivateServerCertificate `json:"privateServerCertificate,omitempty" yaml:"privateServerCertificate,omitempty"`

	/*
	   Enum for Trust Model
	   Possible values are: `PUBLIC`, `PRIVATE`, `INSECURE`.
	*/
	TrustModel string `json:"trustModel,omitempty" yaml:"trustModel,omitempty"`

	/*
	   Enum for controlling the SSL Type (TLS/MTLS)
	   Possible values are: `TLS`, `MTLS`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Bool for enabling SSL
	UseSsl bool `json:"useSsl,omitempty" yaml:"useSsl,omitempty"`

	/*
	   Type of Client Cert (PEM/JKS/.. etc.)
	   Possible values are: `PEM`.
	*/
	ClientCertType string `json:"clientCertType,omitempty" yaml:"clientCertType,omitempty"`

	/*
	   Client Private Key
	   Structure is documented below.
	*/
	ClientPrivateKey Integrationconnectors_ConnectionSslConfigClientPrivateKey `json:"clientPrivateKey,omitempty" yaml:"clientPrivateKey,omitempty"`

	/*
	   Secret containing the passphrase protecting the Client Private Key
	   Structure is documented below.
	*/
	ClientPrivateKeyPass Integrationconnectors_ConnectionSslConfigClientPrivateKeyPass `json:"clientPrivateKeyPass,omitempty" yaml:"clientPrivateKeyPass,omitempty"`
}
