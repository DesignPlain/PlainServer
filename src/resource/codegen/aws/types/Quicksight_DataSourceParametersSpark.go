package types

type Quicksight_DataSourceParametersSpark struct {
	// The warehouse to which to connect.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// The host to which to connect.
	Host string `json:"host,omitempty" yaml:"host,omitempty"`
}
