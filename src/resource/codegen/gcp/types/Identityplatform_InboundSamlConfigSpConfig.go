package types

type Identityplatform_InboundSamlConfigSpConfig struct {
	// Callback URI where responses from IDP are handled. Must start with `https://`.
	CallbackUri string `json:"callbackUri,omitempty" yaml:"callbackUri,omitempty"`

	/*
	   (Output)
	   The IDP's certificate data to verify the signature in the SAMLResponse issued by the IDP.
	   Structure is documented below.


	   <a name="nested_sp_certificates"></a>The `sp_certificates` block contains:
	*/
	SpCertificates []Identityplatform_InboundSamlConfigSpConfigSpCertificate `json:"spCertificates,omitempty" yaml:"spCertificates,omitempty"`

	// Unique identifier for all SAML entities.
	SpEntityId string `json:"spEntityId,omitempty" yaml:"spEntityId,omitempty"`
}
