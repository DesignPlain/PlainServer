package types

type Alloydb_InstanceClientConnectionConfigSslConfig struct {
	/*
	   SSL mode. Specifies client-server SSL/TLS connection behavior.
	   Possible values are: `ENCRYPTED_ONLY`, `ALLOW_UNENCRYPTED_AND_ENCRYPTED`.
	*/
	SslMode string `json:"sslMode,omitempty" yaml:"sslMode,omitempty"`
}
