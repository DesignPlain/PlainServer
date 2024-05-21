package types

type Sql_getDatabaseInstanceServerCaCert struct {
	// The CN valid for the CA Cert.
	CommonName string `json:"commonName,omitempty" yaml:"commonName,omitempty"`

	// Creation time of the CA Cert.
	CreateTime string `json:"createTime,omitempty" yaml:"createTime,omitempty"`

	// Expiration time of the CA Cert.
	ExpirationTime string `json:"expirationTime,omitempty" yaml:"expirationTime,omitempty"`

	// SHA Fingerprint of the CA Cert.
	Sha1Fingerprint string `json:"sha1Fingerprint,omitempty" yaml:"sha1Fingerprint,omitempty"`

	// The CA Certificate used to connect to the SQL Instance via SSL.
	Cert string `json:"cert,omitempty" yaml:"cert,omitempty"`
}
