package dms

import types "libds/aws/types"

type Endpoint struct {
	// ARN for the certificate.
	CertificateArn string `json:"certificateArn,omitempty" yaml:"certificateArn,omitempty"`

	// Type of endpoint. Valid values are `source`, `target`.
	EndpointType string `json:"endpointType,omitempty" yaml:"endpointType,omitempty"`

	// Configuration block for Kafka settings. See below.
	KafkaSettings types.Dms_EndpointKafkaSettings `json:"kafkaSettings,omitempty" yaml:"kafkaSettings,omitempty"`

	/*
	   ARN for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region. To encrypt an S3 target with a KMS Key, use the parameter `s3_settings.server_side_encryption_kms_key_id`. When `engine_name` is `redshift`, `kms_key_arn` is the KMS Key for the Redshift target and the parameter `redshift_settings.server_side_encryption_kms_key_id` encrypts the S3 intermediate storage.

	   The following arguments are optional:
	*/
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// Port used by the endpoint database.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Configuration block for Postgres settings. See below.
	PostgresSettings types.Dms_EndpointPostgresSettings `json:"postgresSettings,omitempty" yaml:"postgresSettings,omitempty"`

	// (--Deprecated--, use the `aws.dms.S3Endpoint` resource instead) Configuration block for S3 settings. See below.
	S3Settings types.Dms_EndpointS3Settings `json:"s3Settings,omitempty" yaml:"s3Settings,omitempty"`

	// SSL mode to use for the connection. Valid values are `none`, `require`, `verify-ca`, `verify-full`
	SslMode string `json:"sslMode,omitempty" yaml:"sslMode,omitempty"`

	// Configuration block for Kinesis settings. See below.
	KinesisSettings types.Dms_EndpointKinesisSettings `json:"kinesisSettings,omitempty" yaml:"kinesisSettings,omitempty"`

	// Configuration block for MongoDB settings. See below.
	MongodbSettings types.Dms_EndpointMongodbSettings `json:"mongodbSettings,omitempty" yaml:"mongodbSettings,omitempty"`

	// Password to be used to login to the endpoint database.
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// Configuration block for Redshift settings. See below.
	RedshiftSettings types.Dms_EndpointRedshiftSettings `json:"redshiftSettings,omitempty" yaml:"redshiftSettings,omitempty"`

	// Full ARN, partial ARN, or friendly name of the Secrets Manager secret that contains the endpoint connection details. Supported only when `engine_name` is `aurora`, `aurora-postgresql`, `mariadb`, `mongodb`, `mysql`, `oracle`, `postgres`, `redshift`, or `sqlserver`.
	SecretsManagerArn string `json:"secretsManagerArn,omitempty" yaml:"secretsManagerArn,omitempty"`

	// User name to be used to login to the endpoint database.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	// Name of the endpoint database.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	//
	PauseReplicationTasks bool `json:"pauseReplicationTasks,omitempty" yaml:"pauseReplicationTasks,omitempty"`

	/*
	   ARN of the IAM role that specifies AWS DMS as the trusted entity and has the required permissions to access the value in the Secrets Manager secret referred to by `secrets_manager_arn`. The role must allow the `iam:PassRole` action.

	   > --Note:-- You can specify one of two sets of values for these permissions. You can specify the values for this setting and `secrets_manager_arn`. Or you can specify clear-text values for `username`, `password` , `server_name`, and `port`. You can't specify both.
	*/
	SecretsManagerAccessRoleArn string `json:"secretsManagerAccessRoleArn,omitempty" yaml:"secretsManagerAccessRoleArn,omitempty"`

	// ARN used by the service access IAM role for dynamodb endpoints.
	ServiceAccessRole string `json:"serviceAccessRole,omitempty" yaml:"serviceAccessRole,omitempty"`

	// Configuration block for OpenSearch settings. See below.
	ElasticsearchSettings types.Dms_EndpointElasticsearchSettings `json:"elasticsearchSettings,omitempty" yaml:"elasticsearchSettings,omitempty"`

	// Database endpoint identifier. Identifiers must contain from 1 to 255 alphanumeric characters or hyphens, begin with a letter, contain only ASCII letters, digits, and hyphens, not end with a hyphen, and not contain two consecutive hyphens.
	EndpointId string `json:"endpointId,omitempty" yaml:"endpointId,omitempty"`

	// Type of engine for the endpoint. Valid values are `aurora`, `aurora-postgresql`, `azuredb`, `azure-sql-managed-instance`, `babelfish`, `db2`, `db2-zos`, `docdb`, `dynamodb`, `elasticsearch`, `kafka`, `kinesis`, `mariadb`, `mongodb`, `mysql`, `opensearch`, `oracle`, `postgres`, `redshift`, `s3`, `sqlserver`, `sybase`. Please note that some of engine names are available only for `target` endpoint type (e.g. `redshift`).
	EngineName string `json:"engineName,omitempty" yaml:"engineName,omitempty"`

	// Additional attributes associated with the connection. For available attributes for a `source` Endpoint, see [Sources for data migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.html). For available attributes for a `target` Endpoint, see [Targets for data migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.html).
	ExtraConnectionAttributes string `json:"extraConnectionAttributes,omitempty" yaml:"extraConnectionAttributes,omitempty"`

	//
	RedisSettings types.Dms_EndpointRedisSettings `json:"redisSettings,omitempty" yaml:"redisSettings,omitempty"`

	// Host name of the server.
	ServerName string `json:"serverName,omitempty" yaml:"serverName,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
