package types

type Integrationconnectors_ConnectionAuthConfigOauth2JwtBearer struct {
	/*
	   Secret version reference containing a PKCS#8 PEM-encoded private key associated with the Client Certificate.
	   This private key will be used to sign JWTs used for the jwt-bearer authorization grant.
	   Specified in the form as: projects/-/secrets/-/versions/-.
	   Structure is documented below.
	*/
	ClientKey Integrationconnectors_ConnectionAuthConfigOauth2JwtBearerClientKey `json:"clientKey,omitempty" yaml:"clientKey,omitempty"`

	/*
	   JwtClaims providers fields to generate the token.
	   Structure is documented below.
	*/
	JwtClaims Integrationconnectors_ConnectionAuthConfigOauth2JwtBearerJwtClaims `json:"jwtClaims,omitempty" yaml:"jwtClaims,omitempty"`
}
