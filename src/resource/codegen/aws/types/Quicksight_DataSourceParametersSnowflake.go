package types

type Quicksight_DataSourceParametersSnowflake struct {
	// The warehouse to which to connect.
	Warehouse string `json:"warehouse,omitempty" yaml:"warehouse,omitempty"`

	// The database to which to connect.
	Database string `json:"database,omitempty" yaml:"database,omitempty"`

	// The host to which to connect.
	Host string `json:"host,omitempty" yaml:"host,omitempty"`
}
