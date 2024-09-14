package types

type Dms_getEndpointRedisSetting struct {
	//
	AuthType string `json:"authType,omitempty" yaml:"authType,omitempty"`

	//
	AuthUserName string `json:"authUserName,omitempty" yaml:"authUserName,omitempty"`

	//
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	//
	ServerName string `json:"serverName,omitempty" yaml:"serverName,omitempty"`

	//
	SslCaCertificateArn string `json:"sslCaCertificateArn,omitempty" yaml:"sslCaCertificateArn,omitempty"`

	//
	SslSecurityProtocol string `json:"sslSecurityProtocol,omitempty" yaml:"sslSecurityProtocol,omitempty"`

	//
	AuthPassword string `json:"authPassword,omitempty" yaml:"authPassword,omitempty"`
}
