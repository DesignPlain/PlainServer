package types

type Certificateauthority_CertificateTemplatePassthroughExtensionsAdditionalExtension struct {
	// Required. The parts of an OID path. The most significant parts of the path come first.
	ObjectIdPaths []int `json:"objectIdPaths,omitempty" yaml:"objectIdPaths,omitempty"`
}
