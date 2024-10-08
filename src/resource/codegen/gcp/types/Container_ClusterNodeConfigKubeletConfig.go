package types

type Container_ClusterNodeConfigKubeletConfig struct {
	// Controls the maximum number of processes allowed to run in a pod. The value must be greater than or equal to 1024 and less than 4194304.
	PodPidsLimit int `json:"podPidsLimit,omitempty" yaml:"podPidsLimit,omitempty"`

	/*
	   If true, enables CPU CFS quota enforcement for
	   containers that specify CPU limits.
	*/
	CpuCfsQuota bool `json:"cpuCfsQuota,omitempty" yaml:"cpuCfsQuota,omitempty"`

	/*
	   The CPU CFS quota period value. Specified
	   as a sequence of decimal numbers, each with optional fraction and a unit suffix,
	   such as `"300ms"`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m",
	   "h". The value must be a positive duration.

	   > Note: At the time of writing (2020/08/18) the GKE API rejects the `none`
	   value and accepts an invalid `default` value instead. While this remains true,
	   not specifying the `kubelet_config` block should be the equivalent of specifying
	   `none`.
	*/
	CpuCfsQuotaPeriod string `json:"cpuCfsQuotaPeriod,omitempty" yaml:"cpuCfsQuotaPeriod,omitempty"`

	/*
	   The CPU management policy on the node. See
	   [K8S CPU Management Policies](https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/).
	   One of `"none"` or `"static"`. Defaults to `none` when `kubelet_config` is unset.
	*/
	CpuManagerPolicy string `json:"cpuManagerPolicy,omitempty" yaml:"cpuManagerPolicy,omitempty"`
}
