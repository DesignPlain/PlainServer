package types

type Appengine_StandardAppVersionAutomaticScaling struct {
	/*
	   Maximum amount of time that a request should wait in the pending queue before starting a new instance to handle it.
	   A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	*/
	MaxPendingLatency string `json:"maxPendingLatency,omitempty" yaml:"maxPendingLatency,omitempty"`

	// Minimum number of idle instances that should be maintained for this version. Only applicable for the default version of a service.
	MinIdleInstances int `json:"minIdleInstances,omitempty" yaml:"minIdleInstances,omitempty"`

	/*
	   Minimum amount of time a request should wait in the pending queue before starting a new instance to handle it.
	   A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	*/
	MinPendingLatency string `json:"minPendingLatency,omitempty" yaml:"minPendingLatency,omitempty"`

	/*
	   Scheduler settings for standard environment.
	   Structure is documented below.
	*/
	StandardSchedulerSettings Appengine_StandardAppVersionAutomaticScalingStandardSchedulerSettings `json:"standardSchedulerSettings,omitempty" yaml:"standardSchedulerSettings,omitempty"`

	/*
	   Number of concurrent requests an automatic scaling instance can accept before the scheduler spawns a new instance.
	   Defaults to a runtime-specific value.
	*/
	MaxConcurrentRequests int `json:"maxConcurrentRequests,omitempty" yaml:"maxConcurrentRequests,omitempty"`

	// Maximum number of idle instances that should be maintained for this version.
	MaxIdleInstances int `json:"maxIdleInstances,omitempty" yaml:"maxIdleInstances,omitempty"`
}
