package types

type Dms_getEndpointKafkaSetting struct {
	//
	IncludeTableAlterOperations bool `json:"includeTableAlterOperations,omitempty" yaml:"includeTableAlterOperations,omitempty"`

	//
	IncludeTransactionDetails bool `json:"includeTransactionDetails,omitempty" yaml:"includeTransactionDetails,omitempty"`

	//
	MessageFormat string `json:"messageFormat,omitempty" yaml:"messageFormat,omitempty"`

	//
	NoHexPrefix bool `json:"noHexPrefix,omitempty" yaml:"noHexPrefix,omitempty"`

	//
	SaslUsername string `json:"saslUsername,omitempty" yaml:"saslUsername,omitempty"`

	//
	SslClientKeyArn string `json:"sslClientKeyArn,omitempty" yaml:"sslClientKeyArn,omitempty"`

	//
	Topic string `json:"topic,omitempty" yaml:"topic,omitempty"`

	//
	Broker string `json:"broker,omitempty" yaml:"broker,omitempty"`

	//
	IncludeControlDetails bool `json:"includeControlDetails,omitempty" yaml:"includeControlDetails,omitempty"`

	//
	SaslPassword string `json:"saslPassword,omitempty" yaml:"saslPassword,omitempty"`

	//
	SecurityProtocol string `json:"securityProtocol,omitempty" yaml:"securityProtocol,omitempty"`

	//
	SslCaCertificateArn string `json:"sslCaCertificateArn,omitempty" yaml:"sslCaCertificateArn,omitempty"`

	//
	SslClientKeyPassword string `json:"sslClientKeyPassword,omitempty" yaml:"sslClientKeyPassword,omitempty"`

	//
	IncludeNullAndEmpty bool `json:"includeNullAndEmpty,omitempty" yaml:"includeNullAndEmpty,omitempty"`

	//
	IncludePartitionValue bool `json:"includePartitionValue,omitempty" yaml:"includePartitionValue,omitempty"`

	//
	MessageMaxBytes int `json:"messageMaxBytes,omitempty" yaml:"messageMaxBytes,omitempty"`

	//
	PartitionIncludeSchemaTable bool `json:"partitionIncludeSchemaTable,omitempty" yaml:"partitionIncludeSchemaTable,omitempty"`

	//
	SslClientCertificateArn string `json:"sslClientCertificateArn,omitempty" yaml:"sslClientCertificateArn,omitempty"`
}
