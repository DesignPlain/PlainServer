package types

type Dataproc_AutoscalingPolicyWorkerConfig struct {
	// Minimum number of instances for this group. Bounds: [2, maxInstances]. Defaults to 2.
	MinInstances int `json:"minInstances,omitempty" yaml:"minInstances,omitempty"`

	/*
	   Weight for the instance group, which is used to determine the fraction of total workers
	   in the cluster from this instance group. For example, if primary workers have weight 2,
	   and secondary workers have weight 1, the cluster will have approximately 2 primary workers
	   for each secondary worker.
	   The cluster may not reach the specified balance if constrained by min/max bounds or other
	   autoscaling settings. For example, if maxInstances for secondary workers is 0, then only
	   primary workers will be added. The cluster can also be out of balance when created.
	   If weight is not set on any instance group, the cluster will default to equal weight for
	   all groups: the cluster will attempt to maintain an equal number of workers in each group
	   within the configured size bounds for each group. If weight is set for one group only,
	   the cluster will default to zero weight on the unset group. For example if weight is set
	   only on primary workers, the cluster will use primary workers only and no secondary workers.
	*/
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`

	// Maximum number of instances for this group.
	MaxInstances int `json:"maxInstances,omitempty" yaml:"maxInstances,omitempty"`
}
