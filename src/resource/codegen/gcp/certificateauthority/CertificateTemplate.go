package certificateauthority

import types "libds/gcp/types"

type CertificateTemplate struct {
	// Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
	IdentityConstraints types.Certificateauthority_CertificateTemplateIdentityConstraints `json:"identityConstraints,omitempty" yaml:"identityConstraints,omitempty"`

	/*
	   Optional. Labels with user-defined metadata.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The resource name for this CertificateTemplate in the format `projects/-/locations/-/certificateTemplates/-`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baseline_values that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
	PassthroughExtensions types.Certificateauthority_CertificateTemplatePassthroughExtensions `json:"passthroughExtensions,omitempty" yaml:"passthroughExtensions,omitempty"`

	// Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baseline_values for the same properties, the certificate issuance request will fail.
	PredefinedValues types.Certificateauthority_CertificateTemplatePredefinedValues `json:"predefinedValues,omitempty" yaml:"predefinedValues,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
