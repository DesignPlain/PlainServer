package types

type Emr_ClusterCoreInstanceFleetLaunchSpecifications struct {
	// Configuration block for on demand instances launch specifications.
	OnDemandSpecifications []Emr_ClusterCoreInstanceFleetLaunchSpecificationsOnDemandSpecification `json:"onDemandSpecifications,omitempty" yaml:"onDemandSpecifications,omitempty"`

	// Configuration block for spot instances launch specifications.
	SpotSpecifications []Emr_ClusterCoreInstanceFleetLaunchSpecificationsSpotSpecification `json:"spotSpecifications,omitempty" yaml:"spotSpecifications,omitempty"`
}
