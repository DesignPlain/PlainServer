package types

type Ssoadmin_TrustedTokenIssuerTrustedTokenIssuerConfigurationOidcJwtConfiguration struct {
	// Specifies the path of the source attribute in the JWT from the trusted token issuer.
	ClaimAttributePath string `json:"claimAttributePath,omitempty" yaml:"claimAttributePath,omitempty"`

	// Specifies path of the destination attribute in a JWT from IAM Identity Center. The attribute mapped by this JMESPath expression is compared against the attribute mapped by `claim_attribute_path` when a trusted token issuer token is exchanged for an IAM Identity Center token.
	IdentityStoreAttributePath string `json:"identityStoreAttributePath,omitempty" yaml:"identityStoreAttributePath,omitempty"`

	// Specifies the URL that IAM Identity Center uses for OpenID Discovery. OpenID Discovery is used to obtain the information required to verify the tokens that the trusted token issuer generates.
	IssuerUrl string `json:"issuerUrl,omitempty" yaml:"issuerUrl,omitempty"`

	// The method that the trusted token issuer can use to retrieve the JSON Web Key Set used to verify a JWT. Valid values are `OPEN_ID_DISCOVERY`
	JwksRetrievalOption string `json:"jwksRetrievalOption,omitempty" yaml:"jwksRetrievalOption,omitempty"`
}
