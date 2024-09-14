package types

type Iam_WorkloadIdentityPoolProviderOidc struct {
	/*
	   Acceptable values for the `aud` field (audience) in the OIDC token. Token exchange
	   requests are rejected if the token audience does not match one of the configured
	   values. Each audience may be at most 256 characters. A maximum of 10 audiences may
	   be configured.
	   If this list is empty, the OIDC token audience must be equal to the full canonical
	   resource name of the WorkloadIdentityPoolProvider, with or without the HTTPS prefix.
	   For example:
	*/
	AllowedAudiences []string `json:"allowedAudiences,omitempty" yaml:"allowedAudiences,omitempty"`

	// The OIDC issuer URL.
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
