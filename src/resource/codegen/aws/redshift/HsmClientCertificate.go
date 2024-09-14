package redshift

type HsmClientCertificate struct {
	// The identifier of the HSM client certificate.
	HsmClientCertificateIdentifier string `json:"hsmClientCertificateIdentifier,omitempty" yaml:"hsmClientCertificateIdentifier,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
