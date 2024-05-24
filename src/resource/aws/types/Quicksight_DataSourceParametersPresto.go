package types

type Quicksight_DataSourceParametersPresto struct {
	// The catalog to which to connect.
	Catalog string `json:"catalog,omitempty" yaml:"catalog,omitempty"`

	// The host to which to connect.
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	// The port to which to connect.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
