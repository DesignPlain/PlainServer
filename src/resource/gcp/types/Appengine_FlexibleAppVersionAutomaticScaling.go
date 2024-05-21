package types

type Appengine_FlexibleAppVersionAutomaticScaling struct {
	// Maximum number of instances that should be started to handle requests for this version. Default: 20
	MaxTotalInstances int `json:"maxTotalInstances,omitempty" yaml:"maxTotalInstances,omitempty"`

	// Minimum amount of time a request should wait in the pending queue before starting a new instance to handle it.
	MinPendingLatency string `json:"minPendingLatency,omitempty" yaml:"minPendingLatency,omitempty"`

	/*
	   Target scaling by CPU usage.
	   Structure is documented below.
	*/
	CpuUtilization Appengine_FlexibleAppVersionAutomaticScalingCpuUtilization `json:"cpuUtilization,omitempty" yaml:"cpuUtilization,omitempty"`

	/*
	   Target scaling by disk usage.
	   Structure is documented below.
	*/
	DiskUtilization Appengine_FlexibleAppVersionAutomaticScalingDiskUtilization `json:"diskUtilization,omitempty" yaml:"diskUtilization,omitempty"`

	/*
	   Number of concurrent requests an automatic scaling instance can accept before the scheduler spawns a new instance.
	   Defaults to a runtime-specific value.
	*/
	MaxConcurrentRequests int `json:"maxConcurrentRequests,omitempty" yaml:"maxConcurrentRequests,omitempty"`

	// Maximum number of idle instances that should be maintained for this version.
	MaxIdleInstances int `json:"maxIdleInstances,omitempty" yaml:"maxIdleInstances,omitempty"`

	// Maximum amount of time that a request should wait in the pending queue before starting a new instance to handle it.
	MaxPendingLatency string `json:"maxPendingLatency,omitempty" yaml:"maxPendingLatency,omitempty"`

	/*
	   The time period that the Autoscaler should wait before it starts collecting information from a new instance.
	   This prevents the autoscaler from collecting information when the instance is initializing,
	   during which the collected usage would not be reliable. Default: 120s
	*/
	CoolDownPeriod string `json:"coolDownPeriod,omitempty" yaml:"coolDownPeriod,omitempty"`

	// Minimum number of idle instances that should be maintained for this version. Only applicable for the default version of a service.
	MinIdleInstances int `json:"minIdleInstances,omitempty" yaml:"minIdleInstances,omitempty"`

	// Minimum number of running instances that should be maintained for this version. Default: 2
	MinTotalInstances int `json:"minTotalInstances,omitempty" yaml:"minTotalInstances,omitempty"`

	/*
	   Target scaling by network usage.
	   Structure is documented below.
	*/
	NetworkUtilization Appengine_FlexibleAppVersionAutomaticScalingNetworkUtilization `json:"networkUtilization,omitempty" yaml:"networkUtilization,omitempty"`

	/*
	   Target scaling by request utilization.
	   Structure is documented below.
	*/
	RequestUtilization Appengine_FlexibleAppVersionAutomaticScalingRequestUtilization `json:"requestUtilization,omitempty" yaml:"requestUtilization,omitempty"`
}
