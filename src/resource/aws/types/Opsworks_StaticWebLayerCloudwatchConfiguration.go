package types

type Opsworks_StaticWebLayerCloudwatchConfiguration struct {
	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	//
	LogStreams []Opsworks_StaticWebLayerCloudwatchConfigurationLogStream `json:"logStreams,omitempty" yaml:"logStreams,omitempty"`
}
