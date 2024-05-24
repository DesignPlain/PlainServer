package types

type Rds_getProxyAuth struct {
	//
	ClientPasswordAuthType string `json:"clientPasswordAuthType,omitempty" yaml:"clientPasswordAuthType,omitempty"`

	//
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	//
	IamAuth string `json:"iamAuth,omitempty" yaml:"iamAuth,omitempty"`

	//
	SecretArn string `json:"secretArn,omitempty" yaml:"secretArn,omitempty"`

	//
	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	//
	AuthScheme string `json:"authScheme,omitempty" yaml:"authScheme,omitempty"`
}
