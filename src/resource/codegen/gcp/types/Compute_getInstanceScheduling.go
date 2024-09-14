package types

type Compute_getInstanceScheduling struct {
	// Specifies the frequency of planned maintenance events. The accepted values are: PERIODIC
	MaintenanceInterval string `json:"maintenanceInterval,omitempty" yaml:"maintenanceInterval,omitempty"`

	// Specifies node affinities or anti-affinities to determine which sole-tenant nodes your instances and managed instance groups will use as host systems.
	NodeAffinities []Compute_getInstanceSchedulingNodeAffinity `json:"nodeAffinities,omitempty" yaml:"nodeAffinities,omitempty"`

	/*
	   Describes maintenance behavior for the
	   instance. One of `MIGRATE` or `TERMINATE`, for more info, read
	   [here](https://cloud.google.com/compute/docs/instances/setting-instance-scheduling-options)
	*/
	OnHostMaintenance string `json:"onHostMaintenance,omitempty" yaml:"onHostMaintenance,omitempty"`

	// Whether the instance is preemptible.
	Preemptible bool `json:"preemptible,omitempty" yaml:"preemptible,omitempty"`

	// Describe the type of preemptible VM.
	ProvisioningModel string `json:"provisioningModel,omitempty" yaml:"provisioningModel,omitempty"`

	/*
	   Specifies if the instance should be
	   restarted if it was terminated by Compute Engine (not a user).
	*/
	AutomaticRestart bool `json:"automaticRestart,omitempty" yaml:"automaticRestart,omitempty"`

	// Describe the type of termination action for `SPOT` VM. Can be `STOP` or `DELETE`.  Read more on [here](https://cloud.google.com/compute/docs/instances/create-use-spot)
	InstanceTerminationAction string `json:"instanceTerminationAction,omitempty" yaml:"instanceTerminationAction,omitempty"`

	/*
	   Specifies the maximum amount of time a Local Ssd Vm should wait while
	     recovery of the Local Ssd state is attempted. Its value should be in
	     between 0 and 168 hours with hour granularity and the default value being 1
	     hour.
	*/
	LocalSsdRecoveryTimeouts []Compute_getInstanceSchedulingLocalSsdRecoveryTimeout `json:"localSsdRecoveryTimeouts,omitempty" yaml:"localSsdRecoveryTimeouts,omitempty"`

	// The timeout for new network connections to hosts.
	MaxRunDurations []Compute_getInstanceSchedulingMaxRunDuration `json:"maxRunDurations,omitempty" yaml:"maxRunDurations,omitempty"`

	//
	MinNodeCpus int `json:"minNodeCpus,omitempty" yaml:"minNodeCpus,omitempty"`
}
