package types

type Identityplatform_InboundSamlConfigIdpConfig struct {
	/*
	   The IdP's certificate data to verify the signature in the SAMLResponse issued by the IDP.
	   Structure is documented below.
	*/
	IdpCertificates []Identityplatform_InboundSamlConfigIdpConfigIdpCertificate `json:"idpCertificates,omitempty" yaml:"idpCertificates,omitempty"`

	// Unique identifier for all SAML entities
	IdpEntityId string `json:"idpEntityId,omitempty" yaml:"idpEntityId,omitempty"`

	// Indicates if outbounding SAMLRequest should be signed.
	SignRequest bool `json:"signRequest,omitempty" yaml:"signRequest,omitempty"`

	// URL to send Authentication request to.
	SsoUrl string `json:"ssoUrl,omitempty" yaml:"ssoUrl,omitempty"`
}
