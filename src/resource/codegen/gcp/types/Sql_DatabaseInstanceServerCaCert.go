package types

type Sql_DatabaseInstanceServerCaCert struct {
	// The CA Certificate used to connect to the SQL Instance via SSL.
	Cert string `json:"cert,omitempty" yaml:"cert,omitempty"`

	// The CN valid for the CA Cert.
	CommonName string `json:"commonName,omitempty" yaml:"commonName,omitempty"`

	// Creation time of the CA Cert.
	CreateTime string `json:"createTime,omitempty" yaml:"createTime,omitempty"`

	/*
	   The [RFC 3339](https://tools.ietf.org/html/rfc3339)
	   formatted date time string indicating when this whitelist expires.
	*/
	ExpirationTime string `json:"expirationTime,omitempty" yaml:"expirationTime,omitempty"`

	// SHA Fingerprint of the CA Cert.
	Sha1Fingerprint string `json:"sha1Fingerprint,omitempty" yaml:"sha1Fingerprint,omitempty"`
}
