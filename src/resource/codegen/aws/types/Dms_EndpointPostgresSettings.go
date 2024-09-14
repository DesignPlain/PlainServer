package types

type Dms_EndpointPostgresSettings struct {
	// Sets the WAL heartbeat frequency (in minutes). Default value is `5`.
	HeartbeatFrequency int `json:"heartbeatFrequency,omitempty" yaml:"heartbeatFrequency,omitempty"`

	// Optional When true, DMS migrates JSONB values as CLOB.
	MapJsonbAsClob bool `json:"mapJsonbAsClob,omitempty" yaml:"mapJsonbAsClob,omitempty"`

	// Optional When true, DMS migrates LONG values as VARCHAR.
	MapLongVarcharAs string `json:"mapLongVarcharAs,omitempty" yaml:"mapLongVarcharAs,omitempty"`

	// Specifies the plugin to use to create a replication slot. Valid values: `pglogical`, `test_decoding`.
	PluginName string `json:"pluginName,omitempty" yaml:"pluginName,omitempty"`

	// Sets the name of a previously created logical replication slot for a CDC load of the PostgreSQL source instance.
	SlotName string `json:"slotName,omitempty" yaml:"slotName,omitempty"`

	// Sets the schema in which the operational DDL database artifacts are created. Default is `public`.
	DdlArtifactsSchema string `json:"ddlArtifactsSchema,omitempty" yaml:"ddlArtifactsSchema,omitempty"`

	// Sets the client statement timeout for the PostgreSQL instance, in seconds. Default value is `60`.
	ExecuteTimeout int `json:"executeTimeout,omitempty" yaml:"executeTimeout,omitempty"`

	// When set to `true`, this value causes a task to fail if the actual size of a LOB column is greater than the specified `LobMaxSize`. Default is `false`.
	FailTasksOnLobTruncation bool `json:"failTasksOnLobTruncation,omitempty" yaml:"failTasksOnLobTruncation,omitempty"`

	// To capture DDL events, AWS DMS creates various artifacts in the PostgreSQL database when the task starts.
	CaptureDdls bool `json:"captureDdls,omitempty" yaml:"captureDdls,omitempty"`

	// You can use PostgreSQL endpoint settings to map a boolean as a boolean from your PostgreSQL source to a Amazon Redshift target. Default value is `false`.
	MapBooleanAsBoolean bool `json:"mapBooleanAsBoolean,omitempty" yaml:"mapBooleanAsBoolean,omitempty"`

	// The Babelfish for Aurora PostgreSQL database name for the endpoint.
	BabelfishDatabaseName string `json:"babelfishDatabaseName,omitempty" yaml:"babelfishDatabaseName,omitempty"`

	// Specifies the default behavior of the replication's handling of PostgreSQL- compatible endpoints that require some additional configuration, such as Babelfish endpoints.
	DatabaseMode string `json:"databaseMode,omitempty" yaml:"databaseMode,omitempty"`

	// Sets the schema in which the heartbeat artifacts are created. Default value is `public`.
	HeartbeatSchema string `json:"heartbeatSchema,omitempty" yaml:"heartbeatSchema,omitempty"`

	// For use with change data capture (CDC) only, this attribute has AWS DMS bypass foreign keys and user triggers to reduce the time it takes to bulk load data.
	AfterConnectScript string `json:"afterConnectScript,omitempty" yaml:"afterConnectScript,omitempty"`

	// The write-ahead log (WAL) heartbeat feature mimics a dummy transaction. By doing this, it prevents idle logical replication slots from holding onto old WAL logs, which can result in storage full situations on the source.
	HeartbeatEnable bool `json:"heartbeatEnable,omitempty" yaml:"heartbeatEnable,omitempty"`

	// Specifies the maximum size (in KB) of any .csv file used to transfer data to PostgreSQL. Default is `32,768 KB`.
	MaxFileSize int `json:"maxFileSize,omitempty" yaml:"maxFileSize,omitempty"`
}
