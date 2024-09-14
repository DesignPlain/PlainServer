package types

type Opsworks_EcsClusterLayerCloudwatchConfiguration struct {
	//
	LogStreams []Opsworks_EcsClusterLayerCloudwatchConfigurationLogStream `json:"logStreams,omitempty" yaml:"logStreams,omitempty"`

	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
