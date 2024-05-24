package types

type Mq_getBrokerEngineTypesBrokerEngineType struct {
	// The MQ engine type to return version details for.
	EngineType string `json:"engineType,omitempty" yaml:"engineType,omitempty"`

	// The list of engine versions.
	EngineVersions []Mq_getBrokerEngineTypesBrokerEngineTypeEngineVersion `json:"engineVersions,omitempty" yaml:"engineVersions,omitempty"`
}
