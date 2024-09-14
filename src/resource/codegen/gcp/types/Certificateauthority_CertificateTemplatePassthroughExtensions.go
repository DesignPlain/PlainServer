package types

type Certificateauthority_CertificateTemplatePassthroughExtensions struct {
	// Optional. A set of named X.509 extensions. Will be combined with additional_extensions to determine the full set of X.509 extensions.
	KnownExtensions []string `json:"knownExtensions,omitempty" yaml:"knownExtensions,omitempty"`

	// Optional. A set of ObjectIds identifying custom X.509 extensions. Will be combined with known_extensions to determine the full set of X.509 extensions.
	AdditionalExtensions []Certificateauthority_CertificateTemplatePassthroughExtensionsAdditionalExtension `json:"additionalExtensions,omitempty" yaml:"additionalExtensions,omitempty"`
}
