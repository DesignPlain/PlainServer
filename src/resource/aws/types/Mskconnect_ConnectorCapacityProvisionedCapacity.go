package types

type Mskconnect_ConnectorCapacityProvisionedCapacity struct {
	// The number of microcontroller units (MCUs) allocated to each connector worker. Valid values: `1`, `2`, `4`, `8`. The default value is `1`.
	McuCount int `json:"mcuCount,omitempty" yaml:"mcuCount,omitempty"`

	// The number of workers that are allocated to the connector.
	WorkerCount int `json:"workerCount,omitempty" yaml:"workerCount,omitempty"`
}
