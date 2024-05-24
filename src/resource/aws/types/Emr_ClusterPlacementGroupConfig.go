package types

type Emr_ClusterPlacementGroupConfig struct {
	// Role of the instance in the cluster. Valid Values: `MASTER`, `CORE`, `TASK`.
	InstanceRole string `json:"instanceRole,omitempty" yaml:"instanceRole,omitempty"`

	// EC2 Placement Group strategy associated with instance role. Valid Values: `SPREAD`, `PARTITION`, `CLUSTER`, `NONE`.
	PlacementStrategy string `json:"placementStrategy,omitempty" yaml:"placementStrategy,omitempty"`
}
