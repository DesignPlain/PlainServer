package types

type Apigee_TargetServerSSlInfo struct {
	// The SSL/TLS cipher suites to be used. For programmable proxies, it must be one of the cipher suite names listed in: http://docs.oracle.com/javase/8/docs/technotes/guides/security/StandardNames.html#ciphersuites. For configurable proxies, it must follow the configuration specified in: https://commondatastorage.googleapis.com/chromium-boringssl-docs/ssl.h.html#Cipher-suite-configuration. This setting has no effect for configurable proxies when negotiating TLS 1.3.
	Ciphers []string `json:"ciphers,omitempty" yaml:"ciphers,omitempty"`

	// Enables two-way TLS.
	ClientAuthEnabled bool `json:"clientAuthEnabled,omitempty" yaml:"clientAuthEnabled,omitempty"`

	// If true, Edge ignores TLS certificate errors. Valid when configuring TLS for target servers and target endpoints, and when configuring virtual hosts that use 2-way TLS. When used with a target endpoint/target server, if the backend system uses SNI and returns a cert with a subject Distinguished Name (DN) that does not match the hostname, there is no way to ignore the error and the connection fails.
	IgnoreValidationErrors bool `json:"ignoreValidationErrors,omitempty" yaml:"ignoreValidationErrors,omitempty"`

	// The resource ID of the truststore.
	TrustStore string `json:"trustStore,omitempty" yaml:"trustStore,omitempty"`

	/*
	   The TLS Common Name of the certificate.
	   Structure is documented below.
	*/
	CommonName Apigee_TargetServerSSlInfoCommonName `json:"commonName,omitempty" yaml:"commonName,omitempty"`

	// Enables TLS. If false, neither one-way nor two-way TLS will be enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Required if clientAuthEnabled is true. The resource ID for the alias containing the private key and cert.
	KeyAlias string `json:"keyAlias,omitempty" yaml:"keyAlias,omitempty"`

	// Required if clientAuthEnabled is true. The resource ID of the keystore.
	KeyStore string `json:"keyStore,omitempty" yaml:"keyStore,omitempty"`

	// The TLS versioins to be used.
	Protocols []string `json:"protocols,omitempty" yaml:"protocols,omitempty"`
}
