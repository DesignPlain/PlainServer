package types

type Emr_InstanceFleetLaunchSpecifications struct {
	// Configuration block for on demand instances launch specifications
	OnDemandSpecifications []Emr_InstanceFleetLaunchSpecificationsOnDemandSpecification `json:"onDemandSpecifications,omitempty" yaml:"onDemandSpecifications,omitempty"`

	// Configuration block for spot instances launch specifications
	SpotSpecifications []Emr_InstanceFleetLaunchSpecificationsSpotSpecification `json:"spotSpecifications,omitempty" yaml:"spotSpecifications,omitempty"`
}
