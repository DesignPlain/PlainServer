package types

type Quicksight_DataSourceParametersRds struct {
	// The database to which to connect.
	Database string `json:"database,omitempty" yaml:"database,omitempty"`

	// The instance ID to which to connect.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`
}
