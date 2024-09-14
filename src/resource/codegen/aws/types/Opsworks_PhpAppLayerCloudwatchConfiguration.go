package types

type Opsworks_PhpAppLayerCloudwatchConfiguration struct {
	//
	LogStreams []Opsworks_PhpAppLayerCloudwatchConfigurationLogStream `json:"logStreams,omitempty" yaml:"logStreams,omitempty"`

	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
