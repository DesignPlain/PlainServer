package rds

import types "DesignSphere_Server/src/resource/aws/types"

type Proxy struct {
	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	RequireTls bool `json:"requireTls,omitempty" yaml:"requireTls,omitempty"`

	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// A mapping of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
	Auths []types.Rds_ProxyAuth `json:"auths,omitempty" yaml:"auths,omitempty"`

	// The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. For Aurora MySQL, RDS for MariaDB, and RDS for MySQL databases, specify `MYSQL`. For Aurora PostgreSQL and RDS for PostgreSQL databases, specify `POSTGRESQL`. For RDS for Microsoft SQL Server, specify `SQLSERVER`. Valid values are `MYSQL`, `POSTGRESQL`, and `SQLSERVER`.
	EngineFamily string `json:"engineFamily,omitempty" yaml:"engineFamily,omitempty"`

	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
	IdleClientTimeout int `json:"idleClientTimeout,omitempty" yaml:"idleClientTimeout,omitempty"`

	// One or more VPC subnet IDs to associate with the new proxy.
	VpcSubnetIds []string `json:"vpcSubnetIds,omitempty" yaml:"vpcSubnetIds,omitempty"`

	// Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
	DebugLogging bool `json:"debugLogging,omitempty" yaml:"debugLogging,omitempty"`

	// The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// One or more VPC security group IDs to associate with the new proxy.
	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" yaml:"vpcSecurityGroupIds,omitempty"`
}
