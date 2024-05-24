package types

type Dms_getEndpointPostgresSetting struct {
	//
	FailTasksOnLobTruncation bool `json:"failTasksOnLobTruncation,omitempty" yaml:"failTasksOnLobTruncation,omitempty"`

	//
	HeartbeatSchema string `json:"heartbeatSchema,omitempty" yaml:"heartbeatSchema,omitempty"`

	//
	MapBooleanAsBoolean bool `json:"mapBooleanAsBoolean,omitempty" yaml:"mapBooleanAsBoolean,omitempty"`

	//
	MapJsonbAsClob bool `json:"mapJsonbAsClob,omitempty" yaml:"mapJsonbAsClob,omitempty"`

	//
	BabelfishDatabaseName string `json:"babelfishDatabaseName,omitempty" yaml:"babelfishDatabaseName,omitempty"`

	//
	CaptureDdls bool `json:"captureDdls,omitempty" yaml:"captureDdls,omitempty"`

	//
	DatabaseMode string `json:"databaseMode,omitempty" yaml:"databaseMode,omitempty"`

	//
	DdlArtifactsSchema string `json:"ddlArtifactsSchema,omitempty" yaml:"ddlArtifactsSchema,omitempty"`

	//
	HeartbeatEnable bool `json:"heartbeatEnable,omitempty" yaml:"heartbeatEnable,omitempty"`

	//
	SlotName string `json:"slotName,omitempty" yaml:"slotName,omitempty"`

	//
	HeartbeatFrequency int `json:"heartbeatFrequency,omitempty" yaml:"heartbeatFrequency,omitempty"`

	//
	MapLongVarcharAs string `json:"mapLongVarcharAs,omitempty" yaml:"mapLongVarcharAs,omitempty"`

	//
	MaxFileSize int `json:"maxFileSize,omitempty" yaml:"maxFileSize,omitempty"`

	//
	AfterConnectScript string `json:"afterConnectScript,omitempty" yaml:"afterConnectScript,omitempty"`

	//
	ExecuteTimeout int `json:"executeTimeout,omitempty" yaml:"executeTimeout,omitempty"`

	//
	PluginName string `json:"pluginName,omitempty" yaml:"pluginName,omitempty"`
}
