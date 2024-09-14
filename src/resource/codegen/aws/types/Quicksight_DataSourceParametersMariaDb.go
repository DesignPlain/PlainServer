package types

type Quicksight_DataSourceParametersMariaDb struct {
	// The host to which to connect.
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	// The port to which to connect.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// The database to which to connect.
	Database string `json:"database,omitempty" yaml:"database,omitempty"`
}
