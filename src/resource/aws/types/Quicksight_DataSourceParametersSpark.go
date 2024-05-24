package types

type Quicksight_DataSourceParametersSpark struct {
	// The host to which to connect.
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	// The warehouse to which to connect.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
