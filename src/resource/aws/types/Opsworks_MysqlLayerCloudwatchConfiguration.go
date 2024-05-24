package types

type Opsworks_MysqlLayerCloudwatchConfiguration struct {
	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	//
	LogStreams []Opsworks_MysqlLayerCloudwatchConfigurationLogStream `json:"logStreams,omitempty" yaml:"logStreams,omitempty"`
}
