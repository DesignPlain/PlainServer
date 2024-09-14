package types

type Sql_getCaCertsCert struct {
	// Creation time of the CA cert.
	CreateTime string `json:"createTime,omitempty" yaml:"createTime,omitempty"`

	// Expiration time of the CA cert.
	ExpirationTime string `json:"expirationTime,omitempty" yaml:"expirationTime,omitempty"`

	// SHA1 fingerprint of the CA cert.
	Sha1Fingerprint string `json:"sha1Fingerprint,omitempty" yaml:"sha1Fingerprint,omitempty"`

	// The CA certificate used to connect to the SQL instance via SSL.
	Cert string `json:"cert,omitempty" yaml:"cert,omitempty"`

	// The CN valid for the CA cert.
	CommonName string `json:"commonName,omitempty" yaml:"commonName,omitempty"`
}
