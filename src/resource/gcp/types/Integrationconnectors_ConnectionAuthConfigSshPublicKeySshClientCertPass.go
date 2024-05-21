package types

type Integrationconnectors_ConnectionAuthConfigSshPublicKeySshClientCertPass struct {
	/*
	   The resource name of the secret version in the format,
	   format as: projects/-/secrets/-/versions/-.

	   <a name="nested_oauth2_auth_code_flow"></a>The `oauth2_auth_code_flow` block supports:
	*/
	SecretVersion string `json:"secretVersion,omitempty" yaml:"secretVersion,omitempty"`
}
