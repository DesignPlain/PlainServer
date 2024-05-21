package types

type Certificateauthority_CertificateTemplatePredefinedValuesAdditionalExtension struct {
	// Optional. Indicates whether or not this extension is critical (i.e., if the client does not know how to handle this extension, the client should consider this to be an error).
	Critical bool `json:"critical,omitempty" yaml:"critical,omitempty"`

	// Required. The OID for this X.509 extension.
	ObjectId Certificateauthority_CertificateTemplatePredefinedValuesAdditionalExtensionObjectId `json:"objectId,omitempty" yaml:"objectId,omitempty"`

	// Required. The value of this X.509 extension.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
