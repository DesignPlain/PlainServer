package types

type Emr_ClusterMasterInstanceFleet struct {
	// Friendly name given to the instance fleet.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	ProvisionedOnDemandCapacity int `json:"provisionedOnDemandCapacity,omitempty" yaml:"provisionedOnDemandCapacity,omitempty"`

	//
	ProvisionedSpotCapacity int `json:"provisionedSpotCapacity,omitempty" yaml:"provisionedSpotCapacity,omitempty"`

	// Target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
	TargetOnDemandCapacity int `json:"targetOnDemandCapacity,omitempty" yaml:"targetOnDemandCapacity,omitempty"`

	// Target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
	TargetSpotCapacity int `json:"targetSpotCapacity,omitempty" yaml:"targetSpotCapacity,omitempty"`

	// ID of the cluster.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Configuration block for instance fleet.
	InstanceTypeConfigs []Emr_ClusterMasterInstanceFleetInstanceTypeConfig `json:"instanceTypeConfigs,omitempty" yaml:"instanceTypeConfigs,omitempty"`

	// Configuration block for launch specification.
	LaunchSpecifications Emr_ClusterMasterInstanceFleetLaunchSpecifications `json:"launchSpecifications,omitempty" yaml:"launchSpecifications,omitempty"`
}
