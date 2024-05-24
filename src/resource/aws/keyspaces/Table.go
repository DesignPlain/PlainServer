package keyspaces

import types "DesignSphere_Server/src/resource/aws/types"

type Table struct {
	// Specifies if point-in-time recovery is enabled or disabled for the table. More information can be found in the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/PointInTimeRecovery.html).
	PointInTimeRecovery types.Keyspaces_TablePointInTimeRecovery `json:"pointInTimeRecovery,omitempty" yaml:"pointInTimeRecovery,omitempty"`

	/*
	   The name of the table.

	   The following arguments are optional:
	*/
	TableName string `json:"tableName,omitempty" yaml:"tableName,omitempty"`

	// Specifies the read/write throughput capacity mode for the table.
	CapacitySpecification types.Keyspaces_TableCapacitySpecification `json:"capacitySpecification,omitempty" yaml:"capacitySpecification,omitempty"`

	// The name of the keyspace that the table is going to be created in.
	KeyspaceName string `json:"keyspaceName,omitempty" yaml:"keyspaceName,omitempty"`

	// The default Time to Live setting in seconds for the table. More information can be found in the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL-how-it-works.html#ttl-howitworks_default_ttl).
	DefaultTimeToLive int `json:"defaultTimeToLive,omitempty" yaml:"defaultTimeToLive,omitempty"`

	// Specifies how the encryption key for encryption at rest is managed for the table. More information can be found in the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html).
	EncryptionSpecification types.Keyspaces_TableEncryptionSpecification `json:"encryptionSpecification,omitempty" yaml:"encryptionSpecification,omitempty"`

	// Describes the schema of the table.
	SchemaDefinition types.Keyspaces_TableSchemaDefinition `json:"schemaDefinition,omitempty" yaml:"schemaDefinition,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Enables Time to Live custom settings for the table. More information can be found in the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL.html).
	Ttl types.Keyspaces_TableTtl `json:"ttl,omitempty" yaml:"ttl,omitempty"`

	// Enables client-side timestamps for the table. By default, the setting is disabled.
	ClientSideTimestamps types.Keyspaces_TableClientSideTimestamps `json:"clientSideTimestamps,omitempty" yaml:"clientSideTimestamps,omitempty"`

	// A description of the table.
	Comment types.Keyspaces_TableComment `json:"comment,omitempty" yaml:"comment,omitempty"`
}
