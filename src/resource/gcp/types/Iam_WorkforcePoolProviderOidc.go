package types

type Iam_WorkforcePoolProviderOidc struct {
	/*
	   Configuration for web single sign-on for the OIDC provider. Here, web sign-in refers to console sign-in and gcloud sign-in through the browser.
	   Structure is documented below.
	*/
	WebSsoConfig Iam_WorkforcePoolProviderOidcWebSsoConfig `json:"webSsoConfig,omitempty" yaml:"webSsoConfig,omitempty"`

	// The client ID. Must match the audience claim of the JWT issued by the identity provider.
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	/*
	   The optional client secret. Required to enable Authorization Code flow for web sign-in.
	   Structure is documented below.
	*/
	ClientSecret Iam_WorkforcePoolProviderOidcClientSecret `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`

	// The OIDC issuer URI. Must be a valid URI using the 'https' scheme.
	IssuerUri string `json:"issuerUri,omitempty" yaml:"issuerUri,omitempty"`

	/*
	   OIDC JWKs in JSON String format. For details on definition of a
	   JWK, see https:tools.ietf.org/html/rfc7517. If not set, then we
	   use the `jwks_uri` from the discovery document fetched from the
	   .well-known path for the `issuer_uri`. Currently, RSA and EC asymmetric
	   keys are supported. The JWK must use following format and include only
	   the following fields:
	*/
	JwksJson string `json:"jwksJson,omitempty" yaml:"jwksJson,omitempty"`
}
