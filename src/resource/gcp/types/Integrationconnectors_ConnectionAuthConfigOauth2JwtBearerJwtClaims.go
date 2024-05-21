package types

type Integrationconnectors_ConnectionAuthConfigOauth2JwtBearerJwtClaims struct {
	/*
	   Value for the "aud" claim.

	   <a name="nested_oauth2_client_credentials"></a>The `oauth2_client_credentials` block supports:
	*/
	Audience string `json:"audience,omitempty" yaml:"audience,omitempty"`

	// Value for the "iss" claim.
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty"`

	// Value for the "sub" claim.
	Subject string `json:"subject,omitempty" yaml:"subject,omitempty"`
}
