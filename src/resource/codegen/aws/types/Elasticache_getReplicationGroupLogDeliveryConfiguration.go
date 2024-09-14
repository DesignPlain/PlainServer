package types

type Elasticache_getReplicationGroupLogDeliveryConfiguration struct {
	//
	LogFormat string `json:"logFormat,omitempty" yaml:"logFormat,omitempty"`

	//
	LogType string `json:"logType,omitempty" yaml:"logType,omitempty"`

	//
	Destination string `json:"destination,omitempty" yaml:"destination,omitempty"`

	//
	DestinationType string `json:"destinationType,omitempty" yaml:"destinationType,omitempty"`
}
