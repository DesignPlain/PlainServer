package types

type Compute_InstanceTemplateScheduling struct {
	/*
	   Defines the maintenance behavior for this
	   instance.
	*/
	OnHostMaintenance string `json:"onHostMaintenance,omitempty" yaml:"onHostMaintenance,omitempty"`

	/*
	   Allows instance to be preempted. This defaults to
	   false. Read more on this
	   [here](https://cloud.google.com/compute/docs/instances/preemptible).
	*/
	Preemptible bool `json:"preemptible,omitempty" yaml:"preemptible,omitempty"`

	// Describe the type of termination action for `SPOT` VM. Can be `STOP` or `DELETE`.  Read more on [here](https://cloud.google.com/compute/docs/instances/create-use-spot)
	InstanceTerminationAction string `json:"instanceTerminationAction,omitempty" yaml:"instanceTerminationAction,omitempty"`

	/*
	   Specifies the maximum amount of time a Local Ssd Vm should wait while
	     recovery of the Local Ssd state is attempted. Its value should be in
	     between 0 and 168 hours with hour granularity and the default value being 1
	     hour.
	*/
	LocalSsdRecoveryTimeouts []Compute_InstanceTemplateSchedulingLocalSsdRecoveryTimeout `json:"localSsdRecoveryTimeouts,omitempty" yaml:"localSsdRecoveryTimeouts,omitempty"`

	// Specifies the frequency of planned maintenance events. The accepted values are: `PERIODIC`.
	MaintenanceInterval string `json:"maintenanceInterval,omitempty" yaml:"maintenanceInterval,omitempty"`

	/*
	   The duration of the instance. Instance will run and be terminated after then, the termination action could be defined in `instance_termination_action`. Only support `DELETE` `instance_termination_action` at this point. Structure is documented below.
	   <a name="nested_max_run_duration"></a>The `max_run_duration` block supports:
	*/
	MaxRunDuration Compute_InstanceTemplateSchedulingMaxRunDuration `json:"maxRunDuration,omitempty" yaml:"maxRunDuration,omitempty"`

	/*
	   Specifies whether the instance should be
	   automatically restarted if it is terminated by Compute Engine (not
	   terminated by a user). This defaults to true.
	*/
	AutomaticRestart bool `json:"automaticRestart,omitempty" yaml:"automaticRestart,omitempty"`

	// Minimum number of cpus for the instance.
	MinNodeCpus int `json:"minNodeCpus,omitempty" yaml:"minNodeCpus,omitempty"`

	/*
	   Specifies node affinities or anti-affinities
	   to determine which sole-tenant nodes your instances and managed instance
	   groups will use as host systems. Read more on sole-tenant node creation
	   [here](https://cloud.google.com/compute/docs/nodes/create-nodes).
	   Structure documented below.
	*/
	NodeAffinities []Compute_InstanceTemplateSchedulingNodeAffinity `json:"nodeAffinities,omitempty" yaml:"nodeAffinities,omitempty"`

	/*
	   Describe the type of preemptible VM. This field accepts the value `STANDARD` or `SPOT`. If the value is `STANDARD`, there will be no discount. If this   is set to `SPOT`,
	   `preemptible` should be `true` and `automatic_restart` should be
	   `false`. For more info about
	   `SPOT`, read [here](https://cloud.google.com/compute/docs/instances/spot)
	*/
	ProvisioningModel string `json:"provisioningModel,omitempty" yaml:"provisioningModel,omitempty"`
}
