package types

type Apigee_TargetServerSSlInfoCommonName struct {
	// The TLS Common Name string of the certificate.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// Indicates whether the cert should be matched against as a wildcard cert.
	WildcardMatch bool `json:"wildcardMatch,omitempty" yaml:"wildcardMatch,omitempty"`
}
