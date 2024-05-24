package types

type Dms_EndpointRedisSettings struct {
	// The username provided with the `auth-role` option of the AuthType setting for a Redis target endpoint.
	AuthUserName string `json:"authUserName,omitempty" yaml:"authUserName,omitempty"`

	// Transmission Control Protocol (TCP) port for the endpoint.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Fully qualified domain name of the endpoint.
	ServerName string `json:"serverName,omitempty" yaml:"serverName,omitempty"`

	// The Amazon Resource Name (ARN) for the certificate authority (CA) that DMS uses to connect to your Redis target endpoint.
	SslCaCertificateArn string `json:"sslCaCertificateArn,omitempty" yaml:"sslCaCertificateArn,omitempty"`

	// The plaintext option doesn't provide Transport Layer Security (TLS) encryption for traffic between endpoint and database. Options include `plaintext`, `ssl-encryption`. The default is `ssl-encryption`.
	SslSecurityProtocol string `json:"sslSecurityProtocol,omitempty" yaml:"sslSecurityProtocol,omitempty"`

	// The password provided with the auth-role and auth-token options of the AuthType setting for a Redis target endpoint.
	AuthPassword string `json:"authPassword,omitempty" yaml:"authPassword,omitempty"`

	// The type of authentication to perform when connecting to a Redis target. Options include `none`, `auth-token`, and `auth-role`. The `auth-token` option requires an `auth_password` value to be provided. The `auth-role` option requires `auth_user_name` and `auth_password` values to be provided.
	AuthType string `json:"authType,omitempty" yaml:"authType,omitempty"`
}
