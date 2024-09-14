package types

type Opsworks_GangliaLayerCloudwatchConfiguration struct {
	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	//
	LogStreams []Opsworks_GangliaLayerCloudwatchConfigurationLogStream `json:"logStreams,omitempty" yaml:"logStreams,omitempty"`
}
