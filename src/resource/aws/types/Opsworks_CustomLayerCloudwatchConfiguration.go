package types

type Opsworks_CustomLayerCloudwatchConfiguration struct {
	// A block the specifies how an opsworks logs look like. See Log Streams.
	LogStreams []Opsworks_CustomLayerCloudwatchConfigurationLogStream `json:"logStreams,omitempty" yaml:"logStreams,omitempty"`

	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
