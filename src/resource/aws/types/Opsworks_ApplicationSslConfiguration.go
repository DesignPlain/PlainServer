package types

type Opsworks_ApplicationSslConfiguration struct {
	// The contents of the certificate's domain.crt file.
	Certificate string `json:"certificate,omitempty" yaml:"certificate,omitempty"`

	// Can be used to specify an intermediate certificate authority key or client authentication.
	Chain string `json:"chain,omitempty" yaml:"chain,omitempty"`

	// The private key; the contents of the certificate's domain.key file.
	PrivateKey string `json:"privateKey,omitempty" yaml:"privateKey,omitempty"`
}
