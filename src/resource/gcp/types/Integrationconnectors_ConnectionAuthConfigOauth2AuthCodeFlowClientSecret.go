package types

type Integrationconnectors_ConnectionAuthConfigOauth2AuthCodeFlowClientSecret struct {
	/*
	   The resource name of the secret version in the format,
	   format as: projects/-/secrets/-/versions/-.
	*/
	SecretVersion string `json:"secretVersion,omitempty" yaml:"secretVersion,omitempty"`
}
