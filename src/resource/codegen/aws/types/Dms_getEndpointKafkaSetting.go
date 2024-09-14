package types

type Dms_getEndpointKafkaSetting struct {
	//
	IncludePartitionValue bool `json:"includePartitionValue,omitempty" yaml:"includePartitionValue,omitempty"`

	//
	IncludeTableAlterOperations bool `json:"includeTableAlterOperations,omitempty" yaml:"includeTableAlterOperations,omitempty"`

	//
	MessageFormat string `json:"messageFormat,omitempty" yaml:"messageFormat,omitempty"`

	//
	SslCaCertificateArn string `json:"sslCaCertificateArn,omitempty" yaml:"sslCaCertificateArn,omitempty"`

	//
	SslClientCertificateArn string `json:"sslClientCertificateArn,omitempty" yaml:"sslClientCertificateArn,omitempty"`

	//
	SslClientKeyPassword string `json:"sslClientKeyPassword,omitempty" yaml:"sslClientKeyPassword,omitempty"`

	//
	IncludeNullAndEmpty bool `json:"includeNullAndEmpty,omitempty" yaml:"includeNullAndEmpty,omitempty"`

	//
	SaslUsername string `json:"saslUsername,omitempty" yaml:"saslUsername,omitempty"`

	//
	Topic string `json:"topic,omitempty" yaml:"topic,omitempty"`

	//
	PartitionIncludeSchemaTable bool `json:"partitionIncludeSchemaTable,omitempty" yaml:"partitionIncludeSchemaTable,omitempty"`

	//
	MessageMaxBytes int `json:"messageMaxBytes,omitempty" yaml:"messageMaxBytes,omitempty"`

	//
	NoHexPrefix bool `json:"noHexPrefix,omitempty" yaml:"noHexPrefix,omitempty"`

	//
	SaslPassword string `json:"saslPassword,omitempty" yaml:"saslPassword,omitempty"`

	//
	SslClientKeyArn string `json:"sslClientKeyArn,omitempty" yaml:"sslClientKeyArn,omitempty"`

	//
	Broker string `json:"broker,omitempty" yaml:"broker,omitempty"`

	//
	IncludeTransactionDetails bool `json:"includeTransactionDetails,omitempty" yaml:"includeTransactionDetails,omitempty"`

	//
	SecurityProtocol string `json:"securityProtocol,omitempty" yaml:"securityProtocol,omitempty"`

	//
	IncludeControlDetails bool `json:"includeControlDetails,omitempty" yaml:"includeControlDetails,omitempty"`
}
