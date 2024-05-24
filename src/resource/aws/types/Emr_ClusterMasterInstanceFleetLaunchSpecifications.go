package types

type Emr_ClusterMasterInstanceFleetLaunchSpecifications struct {
	// Configuration block for spot instances launch specifications.
	SpotSpecifications []Emr_ClusterMasterInstanceFleetLaunchSpecificationsSpotSpecification `json:"spotSpecifications,omitempty" yaml:"spotSpecifications,omitempty"`

	// Configuration block for on demand instances launch specifications.
	OnDemandSpecifications []Emr_ClusterMasterInstanceFleetLaunchSpecificationsOnDemandSpecification `json:"onDemandSpecifications,omitempty" yaml:"onDemandSpecifications,omitempty"`
}
