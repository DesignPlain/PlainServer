package types

type Mq_getInstanceTypeOfferingsBrokerInstanceOption struct {
	// The list of supported deployment modes.
	SupportedDeploymentModes []string `json:"supportedDeploymentModes,omitempty" yaml:"supportedDeploymentModes,omitempty"`

	// The list of supported engine versions.
	SupportedEngineVersions []string `json:"supportedEngineVersions,omitempty" yaml:"supportedEngineVersions,omitempty"`

	// List of available AZs. See Availability Zones. below
	AvailabilityZones []Mq_getInstanceTypeOfferingsBrokerInstanceOptionAvailabilityZone `json:"availabilityZones,omitempty" yaml:"availabilityZones,omitempty"`

	// Filter response by engine type.
	EngineType string `json:"engineType,omitempty" yaml:"engineType,omitempty"`

	// Filter response by host instance type.
	HostInstanceType string `json:"hostInstanceType,omitempty" yaml:"hostInstanceType,omitempty"`

	// Filter response by storage type.
	StorageType string `json:"storageType,omitempty" yaml:"storageType,omitempty"`
}
