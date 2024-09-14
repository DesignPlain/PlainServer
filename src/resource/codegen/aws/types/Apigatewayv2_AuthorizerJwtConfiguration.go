package types

type Apigatewayv2_AuthorizerJwtConfiguration struct {
	// List of the intended recipients of the JWT. A valid JWT must provide an aud that matches at least one entry in this list.
	Audiences []string `json:"audiences,omitempty" yaml:"audiences,omitempty"`

	// Base domain of the identity provider that issues JSON Web Tokens, such as the `endpoint` attribute of the `aws.cognito.UserPool` resource.
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty"`
}
