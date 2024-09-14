package dynamodb

import types "libds/aws/types"

type Table struct {
	// ARN of the source table to restore. Must be supplied for cross-region restores.
	RestoreSourceTableArn string `json:"restoreSourceTableArn,omitempty" yaml:"restoreSourceTableArn,omitempty"`

	// If set, restores table to the most recent point-in-time recovery point.
	RestoreToLatestTime bool `json:"restoreToLatestTime,omitempty" yaml:"restoreToLatestTime,omitempty"`

	// When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
	StreamViewType string `json:"streamViewType,omitempty" yaml:"streamViewType,omitempty"`

	// Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
	BillingMode string `json:"billingMode,omitempty" yaml:"billingMode,omitempty"`

	// Number of read units for this table. If the `billing_mode` is `PROVISIONED`, this field is required.
	ReadCapacity int `json:"readCapacity,omitempty" yaml:"readCapacity,omitempty"`

	// Name of the table to restore. Must match the name of an existing table.
	RestoreSourceName string `json:"restoreSourceName,omitempty" yaml:"restoreSourceName,omitempty"`

	// Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS-owned Customer Master Key if this argument isn't specified. Must be supplied for cross-region restores. See below.
	ServerSideEncryption types.Dynamodb_TableServerSideEncryption `json:"serverSideEncryption,omitempty" yaml:"serverSideEncryption,omitempty"`

	// A map of tags to populate on the created table. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Describe a GSI for the table; subject to the normal limits on the number of GSIs, projected attributes, etc. See below.
	GlobalSecondaryIndexes []types.Dynamodb_TableGlobalSecondaryIndex `json:"globalSecondaryIndexes,omitempty" yaml:"globalSecondaryIndexes,omitempty"`

	// Time of the point-in-time recovery point to restore.
	RestoreDateTime string `json:"restoreDateTime,omitempty" yaml:"restoreDateTime,omitempty"`

	// Describe an LSI on the table; these can only be allocated _at creation_ so you cannot change this definition after you have created the resource. See below.
	LocalSecondaryIndexes []types.Dynamodb_TableLocalSecondaryIndex `json:"localSecondaryIndexes,omitempty" yaml:"localSecondaryIndexes,omitempty"`

	// Configuration block(s) with [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html) replication configurations. See below.
	Replicas []types.Dynamodb_TableReplica `json:"replicas,omitempty" yaml:"replicas,omitempty"`

	/*
	   Storage class of the table.
	   Valid values are `STANDARD` and `STANDARD_INFREQUENT_ACCESS`.
	   Default value is `STANDARD`.
	*/
	TableClass string `json:"tableClass,omitempty" yaml:"tableClass,omitempty"`

	// Number of write units for this table. If the `billing_mode` is `PROVISIONED`, this field is required.
	WriteCapacity int `json:"writeCapacity,omitempty" yaml:"writeCapacity,omitempty"`

	// Set of nested attribute definitions. Only required for `hash_key` and `range_key` attributes. See below.
	Attributes []types.Dynamodb_TableAttribute `json:"attributes,omitempty" yaml:"attributes,omitempty"`

	// Attribute to use as the hash (partition) key. Must also be defined as an `attribute`. See below.
	HashKey string `json:"hashKey,omitempty" yaml:"hashKey,omitempty"`

	/*
	   Unique within a region name of the table.

	   Optional arguments:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Enable point-in-time recovery options. See below.
	PointInTimeRecovery types.Dynamodb_TablePointInTimeRecovery `json:"pointInTimeRecovery,omitempty" yaml:"pointInTimeRecovery,omitempty"`

	// Attribute to use as the range (sort) key. Must also be defined as an `attribute`, see below.
	RangeKey string `json:"rangeKey,omitempty" yaml:"rangeKey,omitempty"`

	// Whether Streams are enabled.
	StreamEnabled bool `json:"streamEnabled,omitempty" yaml:"streamEnabled,omitempty"`

	// Configuration block for TTL. See below.
	Ttl types.Dynamodb_TableTtl `json:"ttl,omitempty" yaml:"ttl,omitempty"`

	// Enables deletion protection for table. Defaults to `false`.
	DeletionProtectionEnabled bool `json:"deletionProtectionEnabled,omitempty" yaml:"deletionProtectionEnabled,omitempty"`

	// Import Amazon S3 data into a new table. See below.
	ImportTable types.Dynamodb_TableImportTable `json:"importTable,omitempty" yaml:"importTable,omitempty"`
}
