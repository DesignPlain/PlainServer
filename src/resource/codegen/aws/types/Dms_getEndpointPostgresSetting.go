package types

type Dms_getEndpointPostgresSetting struct {
	//
	CaptureDdls bool `json:"captureDdls,omitempty" yaml:"captureDdls,omitempty"`

	//
	DdlArtifactsSchema string `json:"ddlArtifactsSchema,omitempty" yaml:"ddlArtifactsSchema,omitempty"`

	//
	MapJsonbAsClob bool `json:"mapJsonbAsClob,omitempty" yaml:"mapJsonbAsClob,omitempty"`

	//
	BabelfishDatabaseName string `json:"babelfishDatabaseName,omitempty" yaml:"babelfishDatabaseName,omitempty"`

	//
	DatabaseMode string `json:"databaseMode,omitempty" yaml:"databaseMode,omitempty"`

	//
	ExecuteTimeout int `json:"executeTimeout,omitempty" yaml:"executeTimeout,omitempty"`

	//
	FailTasksOnLobTruncation bool `json:"failTasksOnLobTruncation,omitempty" yaml:"failTasksOnLobTruncation,omitempty"`

	//
	HeartbeatFrequency int `json:"heartbeatFrequency,omitempty" yaml:"heartbeatFrequency,omitempty"`

	//
	HeartbeatSchema string `json:"heartbeatSchema,omitempty" yaml:"heartbeatSchema,omitempty"`

	//
	MapLongVarcharAs string `json:"mapLongVarcharAs,omitempty" yaml:"mapLongVarcharAs,omitempty"`

	//
	MaxFileSize int `json:"maxFileSize,omitempty" yaml:"maxFileSize,omitempty"`

	//
	PluginName string `json:"pluginName,omitempty" yaml:"pluginName,omitempty"`

	//
	SlotName string `json:"slotName,omitempty" yaml:"slotName,omitempty"`

	//
	AfterConnectScript string `json:"afterConnectScript,omitempty" yaml:"afterConnectScript,omitempty"`

	//
	HeartbeatEnable bool `json:"heartbeatEnable,omitempty" yaml:"heartbeatEnable,omitempty"`

	//
	MapBooleanAsBoolean bool `json:"mapBooleanAsBoolean,omitempty" yaml:"mapBooleanAsBoolean,omitempty"`
}
