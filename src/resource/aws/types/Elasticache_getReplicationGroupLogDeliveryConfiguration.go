package types

type Elasticache_getReplicationGroupLogDeliveryConfiguration struct {
	//
	Destination string `json:"destination,omitempty" yaml:"destination,omitempty"`

	//
	DestinationType string `json:"destinationType,omitempty" yaml:"destinationType,omitempty"`

	//
	LogFormat string `json:"logFormat,omitempty" yaml:"logFormat,omitempty"`

	//
	LogType string `json:"logType,omitempty" yaml:"logType,omitempty"`
}
