package transfer

type Certificate struct {
	// A short description that helps identify the certificate.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The private key associated with the certificate being imported.
	PrivateKey string `json:"privateKey,omitempty" yaml:"privateKey,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies if a certificate is being used for signing or encryption. The valid values are SIGNING and ENCRYPTION.
	Usage string `json:"usage,omitempty" yaml:"usage,omitempty"`

	// The valid certificate file required for the transfer.
	Certificate string `json:"certificate,omitempty" yaml:"certificate,omitempty"`

	// The optional list of certificate that make up the chain for the certificate that is being imported.
	CertificateChain string `json:"certificateChain,omitempty" yaml:"certificateChain,omitempty"`
}
