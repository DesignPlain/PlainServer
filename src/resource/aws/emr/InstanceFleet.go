package emr

import types "DesignSphere_Server/src/resource/aws/types"

type InstanceFleet struct {
	// ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
	ClusterId string `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`

	// Configuration block for instance fleet
	InstanceTypeConfigs []types.Emr_InstanceFleetInstanceTypeConfig `json:"instanceTypeConfigs,omitempty" yaml:"instanceTypeConfigs,omitempty"`

	// Configuration block for launch specification
	LaunchSpecifications types.Emr_InstanceFleetLaunchSpecifications `json:"launchSpecifications,omitempty" yaml:"launchSpecifications,omitempty"`

	// Friendly name given to the instance fleet.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
	TargetOnDemandCapacity int `json:"targetOnDemandCapacity,omitempty" yaml:"targetOnDemandCapacity,omitempty"`

	// The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
	TargetSpotCapacity int `json:"targetSpotCapacity,omitempty" yaml:"targetSpotCapacity,omitempty"`
}
