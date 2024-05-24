package types

type Opsworks_HaproxyLayerCloudwatchConfiguration struct {
	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	//
	LogStreams []Opsworks_HaproxyLayerCloudwatchConfigurationLogStream `json:"logStreams,omitempty" yaml:"logStreams,omitempty"`
}
