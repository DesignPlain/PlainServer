package types

type Mskconnect_ConnectorCapacity struct {
	// Information about the auto scaling parameters for the connector. See below.
	Autoscaling Mskconnect_ConnectorCapacityAutoscaling `json:"autoscaling,omitempty" yaml:"autoscaling,omitempty"`

	// Details about a fixed capacity allocated to a connector. See below.
	ProvisionedCapacity Mskconnect_ConnectorCapacityProvisionedCapacity `json:"provisionedCapacity,omitempty" yaml:"provisionedCapacity,omitempty"`
}
