package types

type Rds_ProxyAuth struct {
	// Whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy. One of `DISABLED`, `REQUIRED`.
	IamAuth string `json:"iamAuth,omitempty" yaml:"iamAuth,omitempty"`

	// The Amazon Resource Name (ARN) representing the secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster. These secrets are stored within Amazon Secrets Manager.
	SecretArn string `json:"secretArn,omitempty" yaml:"secretArn,omitempty"`

	// The name of the database user to which the proxy connects.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	// The type of authentication that the proxy uses for connections from the proxy to the underlying database. One of `SECRETS`.
	AuthScheme string `json:"authScheme,omitempty" yaml:"authScheme,omitempty"`

	// The type of authentication the proxy uses for connections from clients. Valid values are `MYSQL_NATIVE_PASSWORD`, `POSTGRES_SCRAM_SHA_256`, `POSTGRES_MD5`, and `SQL_SERVER_AUTHENTICATION`.
	ClientPasswordAuthType string `json:"clientPasswordAuthType,omitempty" yaml:"clientPasswordAuthType,omitempty"`

	// A user-specified description about the authentication used by a proxy to log in as a specific database user.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
