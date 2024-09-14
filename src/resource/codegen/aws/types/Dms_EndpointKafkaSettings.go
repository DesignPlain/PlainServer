package types

type Dms_EndpointKafkaSettings struct {
	// Provides detailed transaction information from the source database. This information includes a commit timestamp, a log position, and values for `transaction_id`, previous `transaction_id`, and `transaction_record_id` (the record offset within a transaction). Default is `false`.
	IncludeTransactionDetails bool `json:"includeTransactionDetails,omitempty" yaml:"includeTransactionDetails,omitempty"`

	// Prefixes schema and table names to partition values, when the partition type is `primary-key-type`. Doing this increases data distribution among Kafka partitions. For example, suppose that a SysBench schema has thousands of tables and each table has only limited range for a primary key. In this case, the same primary key is sent from thousands of tables to the same partition, which causes throttling. Default is `false`.
	PartitionIncludeSchemaTable bool `json:"partitionIncludeSchemaTable,omitempty" yaml:"partitionIncludeSchemaTable,omitempty"`

	// Password for the client private key used to securely connect to a Kafka target endpoint.
	SslClientKeyPassword string `json:"sslClientKeyPassword,omitempty" yaml:"sslClientKeyPassword,omitempty"`

	// Kafka topic for migration. Default is `kafka-default-topic`.
	Topic string `json:"topic,omitempty" yaml:"topic,omitempty"`

	// Include NULL and empty columns for records migrated to the endpoint. Default is `false`.
	IncludeNullAndEmpty bool `json:"includeNullAndEmpty,omitempty" yaml:"includeNullAndEmpty,omitempty"`

	// Output format for the records created on the endpoint. Message format is `JSON` (default) or `JSON_UNFORMATTED` (a single line with no tab).
	MessageFormat string `json:"messageFormat,omitempty" yaml:"messageFormat,omitempty"`

	// Maximum size in bytes for records created on the endpoint Default is `1,000,000`.
	MessageMaxBytes int `json:"messageMaxBytes,omitempty" yaml:"messageMaxBytes,omitempty"`

	// Set this optional parameter to true to avoid adding a '0x' prefix to raw data in hexadecimal format. For example, by default, AWS DMS adds a '0x' prefix to the LOB column type in hexadecimal format moving from an Oracle source to a Kafka target. Use the `no_hex_prefix` endpoint setting to enable migration of RAW data type columns without adding the `'0x'` prefix.
	NoHexPrefix bool `json:"noHexPrefix,omitempty" yaml:"noHexPrefix,omitempty"`

	// ARN for the private certificate authority (CA) cert that AWS DMS uses to securely connect to your Kafka target endpoint.
	SslCaCertificateArn string `json:"sslCaCertificateArn,omitempty" yaml:"sslCaCertificateArn,omitempty"`

	// ARN for the client private key used to securely connect to a Kafka target endpoint.
	SslClientKeyArn string `json:"sslClientKeyArn,omitempty" yaml:"sslClientKeyArn,omitempty"`

	// Set secure connection to a Kafka target endpoint using Transport Layer Security (TLS). Options include `ssl-encryption`, `ssl-authentication`, and `sasl-ssl`. `sasl-ssl` requires `sasl_username` and `sasl_password`.
	SecurityProtocol string `json:"securityProtocol,omitempty" yaml:"securityProtocol,omitempty"`

	// ARN of the client certificate used to securely connect to a Kafka target endpoint.
	SslClientCertificateArn string `json:"sslClientCertificateArn,omitempty" yaml:"sslClientCertificateArn,omitempty"`

	// Kafka broker location. Specify in the form broker-hostname-or-ip:port.
	Broker string `json:"broker,omitempty" yaml:"broker,omitempty"`

	// Shows detailed control information for table definition, column definition, and table and column changes in the Kafka message output. Default is `false`.
	IncludeControlDetails bool `json:"includeControlDetails,omitempty" yaml:"includeControlDetails,omitempty"`

	// Shows the partition value within the Kafka message output unless the partition type is `schema-table-type`. Default is `false`.
	IncludePartitionValue bool `json:"includePartitionValue,omitempty" yaml:"includePartitionValue,omitempty"`

	// Includes any data definition language (DDL) operations that change the table in the control data, such as `rename-table`, `drop-table`, `add-column`, `drop-column`, and `rename-column`. Default is `false`.
	IncludeTableAlterOperations bool `json:"includeTableAlterOperations,omitempty" yaml:"includeTableAlterOperations,omitempty"`

	// Secure password you created when you first set up your MSK cluster to validate a client identity and make an encrypted connection between server and client using SASL-SSL authentication.
	SaslPassword string `json:"saslPassword,omitempty" yaml:"saslPassword,omitempty"`

	// Secure user name you created when you first set up your MSK cluster to validate a client identity and make an encrypted connection between server and client using SASL-SSL authentication.
	SaslUsername string `json:"saslUsername,omitempty" yaml:"saslUsername,omitempty"`
}
