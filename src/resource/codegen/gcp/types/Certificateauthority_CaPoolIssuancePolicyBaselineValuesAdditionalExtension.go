package types

type Certificateauthority_CaPoolIssuancePolicyBaselineValuesAdditionalExtension struct {
	/*
	   Indicates whether or not this extension is critical (i.e., if the client does not know how to
	   handle this extension, the client should consider this to be an error).
	*/
	Critical bool `json:"critical,omitempty" yaml:"critical,omitempty"`

	/*
	   Describes values that are relevant in a CA certificate.
	   Structure is documented below.
	*/
	ObjectId Certificateauthority_CaPoolIssuancePolicyBaselineValuesAdditionalExtensionObjectId `json:"objectId,omitempty" yaml:"objectId,omitempty"`

	// The value of this X.509 extension. A base64-encoded string.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
