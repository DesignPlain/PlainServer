package types

type Databasemigrationservice_ConnectionProfilePostgresqlSsl struct {
	/*
	   Input only. The unencrypted PKCS#1 or PKCS#8 PEM-encoded private key associated with the Client Certificate.
	   If this field is used then the 'clientCertificate' field is mandatory.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	ClientKey string `json:"clientKey,omitempty" yaml:"clientKey,omitempty"`

	/*
	   (Output)
	   The current connection profile state.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   Required. Input only. The x509 PEM-encoded certificate of the CA that signed the source database server's certificate.
	   The replica will use this certificate to verify it's connecting to the right host.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	CaCertificate string `json:"caCertificate,omitempty" yaml:"caCertificate,omitempty"`

	/*
	   Input only. The x509 PEM-encoded certificate that will be used by the replica to authenticate against the source database server.
	   If this field is used then the 'clientKey' field is mandatory
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	ClientCertificate string `json:"clientCertificate,omitempty" yaml:"clientCertificate,omitempty"`
}
