package types

type Datastream_ConnectionProfileMysqlProfileSslConfig struct {
	/*
	   PEM-encoded certificate of the CA that signed the source database
	   server's certificate.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	CaCertificate string `json:"caCertificate,omitempty" yaml:"caCertificate,omitempty"`

	/*
	   (Output)
	   Indicates whether the clientKey field is set.
	*/
	CaCertificateSet bool `json:"caCertificateSet,omitempty" yaml:"caCertificateSet,omitempty"`

	/*
	   PEM-encoded certificate that will be used by the replica to
	   authenticate against the source database server. If this field
	   is used then the 'clientKey' and the 'caCertificate' fields are
	   mandatory.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	ClientCertificate string `json:"clientCertificate,omitempty" yaml:"clientCertificate,omitempty"`

	/*
	   (Output)
	   Indicates whether the clientCertificate field is set.
	*/
	ClientCertificateSet bool `json:"clientCertificateSet,omitempty" yaml:"clientCertificateSet,omitempty"`

	/*
	   PEM-encoded private key associated with the Client Certificate.
	   If this field is used then the 'client_certificate' and the
	   'ca_certificate' fields are mandatory.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	ClientKey string `json:"clientKey,omitempty" yaml:"clientKey,omitempty"`

	/*
	   (Output)
	   Indicates whether the clientKey field is set.
	*/
	ClientKeySet bool `json:"clientKeySet,omitempty" yaml:"clientKeySet,omitempty"`
}
