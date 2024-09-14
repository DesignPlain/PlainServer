package types

type Finspace_KxClusterScalingGroupConfiguration struct {
	// A unique identifier for the kdb scaling group.
	ScalingGroupName string `json:"scalingGroupName,omitempty" yaml:"scalingGroupName,omitempty"`

	// The number of vCPUs that you want to reserve for each node of this kdb cluster on the scaling group host.
	Cpu float64 `json:"cpu,omitempty" yaml:"cpu,omitempty"`

	// An optional hard limit on the amount of memory a kdb cluster can use.
	MemoryLimit int `json:"memoryLimit,omitempty" yaml:"memoryLimit,omitempty"`

	// A reservation of the minimum amount of memory that should be available on the scaling group for a kdb cluster to be successfully placed in a scaling group.
	MemoryReservation int `json:"memoryReservation,omitempty" yaml:"memoryReservation,omitempty"`

	// The number of kdb cluster nodes.
	NodeCount int `json:"nodeCount,omitempty" yaml:"nodeCount,omitempty"`
}
