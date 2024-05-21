package types

type Compute_InstanceFromMachineImageScheduling struct {
	// Specifies the action GCE should take when SPOT VM is preempted.
	InstanceTerminationAction string `json:"instanceTerminationAction,omitempty" yaml:"instanceTerminationAction,omitempty"`

	/*
	   Specifies the maximum amount of time a Local Ssd Vm should wait while
	     recovery of the Local Ssd state is attempted. Its value should be in
	     between 0 and 168 hours with hour granularity and the default value being 1
	     hour.
	*/
	LocalSsdRecoveryTimeout Compute_InstanceFromMachineImageSchedulingLocalSsdRecoveryTimeout `json:"localSsdRecoveryTimeout,omitempty" yaml:"localSsdRecoveryTimeout,omitempty"`

	// Specifies the frequency of planned maintenance events. The accepted values are: PERIODIC
	MaintenanceInterval string `json:"maintenanceInterval,omitempty" yaml:"maintenanceInterval,omitempty"`

	//
	MinNodeCpus int `json:"minNodeCpus,omitempty" yaml:"minNodeCpus,omitempty"`

	// Whether the instance is preemptible.
	Preemptible bool `json:"preemptible,omitempty" yaml:"preemptible,omitempty"`

	// Specifies if the instance should be restarted if it was terminated by Compute Engine (not a user).
	AutomaticRestart bool `json:"automaticRestart,omitempty" yaml:"automaticRestart,omitempty"`

	// The timeout for new network connections to hosts.
	MaxRunDuration Compute_InstanceFromMachineImageSchedulingMaxRunDuration `json:"maxRunDuration,omitempty" yaml:"maxRunDuration,omitempty"`

	// Specifies node affinities or anti-affinities to determine which sole-tenant nodes your instances and managed instance groups will use as host systems.
	NodeAffinities []Compute_InstanceFromMachineImageSchedulingNodeAffinity `json:"nodeAffinities,omitempty" yaml:"nodeAffinities,omitempty"`

	// Describes maintenance behavior for the instance. One of MIGRATE or TERMINATE,
	OnHostMaintenance string `json:"onHostMaintenance,omitempty" yaml:"onHostMaintenance,omitempty"`

	// Whether the instance is spot. If this is set as SPOT.
	ProvisioningModel string `json:"provisioningModel,omitempty" yaml:"provisioningModel,omitempty"`
}
