package compute

import types "libds/gcp/types"

type ResourcePolicy struct {
	/*
	   Resource policy for instances used for placement configuration.
	   Structure is documented below.
	*/
	GroupPlacementPolicy types.Compute_ResourcePolicyGroupPlacementPolicy `json:"groupPlacementPolicy,omitempty" yaml:"groupPlacementPolicy,omitempty"`

	/*
	   Resource policy for scheduling instance operations.
	   Structure is documented below.
	*/
	InstanceSchedulePolicy types.Compute_ResourcePolicyInstanceSchedulePolicy `json:"instanceSchedulePolicy,omitempty" yaml:"instanceSchedulePolicy,omitempty"`

	/*
	   The name of the resource, provided by the client when initially creating
	   the resource. The resource name must be 1-63 characters long, and comply
	   with RFC1035. Specifically, the name must be 1-63 characters long and
	   match the regular expression `a-z`? which means the
	   first character must be a lowercase letter, and all following characters
	   must be a dash, lowercase letter, or digit, except the last character,
	   which cannot be a dash.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Region where resource policy resides.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   Policy for creating snapshots of persistent disks.
	   Structure is documented below.
	*/
	SnapshotSchedulePolicy types.Compute_ResourcePolicySnapshotSchedulePolicy `json:"snapshotSchedulePolicy,omitempty" yaml:"snapshotSchedulePolicy,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Replication consistency group for asynchronous disk replication.
	   Structure is documented below.
	*/
	DiskConsistencyGroupPolicy types.Compute_ResourcePolicyDiskConsistencyGroupPolicy `json:"diskConsistencyGroupPolicy,omitempty" yaml:"diskConsistencyGroupPolicy,omitempty"`
}
