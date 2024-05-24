package types

type Opsworks_RailsAppLayerCloudwatchConfiguration struct {
	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	//
	LogStreams []Opsworks_RailsAppLayerCloudwatchConfigurationLogStream `json:"logStreams,omitempty" yaml:"logStreams,omitempty"`
}
