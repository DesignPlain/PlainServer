package types

type Osconfig_PatchDeploymentInstanceFilter struct {
	/*
	   Targets VMs whose name starts with one of these prefixes. Similar to labels, this is another way to group
	   VMs when targeting configs, for example prefix="prod-".
	*/
	InstanceNamePrefixes []string `json:"instanceNamePrefixes,omitempty" yaml:"instanceNamePrefixes,omitempty"`

	/*
	   Targets any of the VM instances specified. Instances are specified by their URI in the `form zones/{{zone}}/instances/{{instance_name}}`,
	   `projects/{{project_id}}/zones/{{zone}}/instances/{{instance_name}}`, or
	   `https://www.googleapis.com/compute/v1/projects/{{project_id}}/zones/{{zone}}/instances/{{instance_name}}`
	*/
	Instances []string `json:"instances,omitempty" yaml:"instances,omitempty"`

	// Targets VM instances in ANY of these zones. Leave empty to target VM instances in any zone.
	Zones []string `json:"zones,omitempty" yaml:"zones,omitempty"`

	// Target all VM instances in the project. If true, no other criteria is permitted.
	All bool `json:"all,omitempty" yaml:"all,omitempty"`

	/*
	   Targets VM instances matching ANY of these GroupLabels. This allows targeting of disparate groups of VM instances.
	   Structure is documented below.
	*/
	GroupLabels []Osconfig_PatchDeploymentInstanceFilterGroupLabel `json:"groupLabels,omitempty" yaml:"groupLabels,omitempty"`
}
