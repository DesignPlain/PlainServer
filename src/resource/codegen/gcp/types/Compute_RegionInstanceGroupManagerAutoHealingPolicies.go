package types

type Compute_RegionInstanceGroupManagerAutoHealingPolicies struct {
	// The health check resource that signals autohealing.
	HealthCheck string `json:"healthCheck,omitempty" yaml:"healthCheck,omitempty"`

	/*
	   The number of seconds that the managed instance group waits before
	   it applies autohealing policies to new instances or recently recreated instances. Between 0 and 3600.
	*/
	InitialDelaySec int `json:"initialDelaySec,omitempty" yaml:"initialDelaySec,omitempty"`
}
