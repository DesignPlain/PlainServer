package types

type Mskconnect_ConnectorCapacityAutoscaling struct {
	// The number of microcontroller units (MCUs) allocated to each connector worker. Valid values: `1`, `2`, `4`, `8`. The default value is `1`.
	McuCount int `json:"mcuCount,omitempty" yaml:"mcuCount,omitempty"`

	// The minimum number of workers allocated to the connector.
	MinWorkerCount int `json:"minWorkerCount,omitempty" yaml:"minWorkerCount,omitempty"`

	// The scale-in policy for the connector. See below.
	ScaleInPolicy Mskconnect_ConnectorCapacityAutoscalingScaleInPolicy `json:"scaleInPolicy,omitempty" yaml:"scaleInPolicy,omitempty"`

	// The scale-out policy for the connector. See below.
	ScaleOutPolicy Mskconnect_ConnectorCapacityAutoscalingScaleOutPolicy `json:"scaleOutPolicy,omitempty" yaml:"scaleOutPolicy,omitempty"`

	// The maximum number of workers allocated to the connector.
	MaxWorkerCount int `json:"maxWorkerCount,omitempty" yaml:"maxWorkerCount,omitempty"`
}
