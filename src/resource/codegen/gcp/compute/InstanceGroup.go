package compute

import types "libds/gcp/types"

type InstanceGroup struct {
	/*
	   The ID of the project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The zone that this instance group should be created in.

	   - - -
	*/
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`

	/*
	   An optional textual description of the instance
	   group.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The list of instances in the group, in `self_link` format.
	   When adding instances they must all be in the same network and zone as the instance group.
	*/
	Instances []string `json:"instances,omitempty" yaml:"instances,omitempty"`

	/*
	   The name of the instance group. Must be 1-63
	   characters long and comply with
	   [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
	   include lowercase letters, numbers, and hyphens.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The named port configuration. See the section below
	   for details on configuration. Structure is documented below.
	*/
	NamedPorts []types.Compute_InstanceGroupNamedPort `json:"namedPorts,omitempty" yaml:"namedPorts,omitempty"`

	/*
	   The URL of the network the instance group is in. If
	   this is different from the network where the instances are in, the creation
	   fails. Defaults to the network where the instances are in (if neither
	   `network` nor `instances` is specified, this field will be blank).
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`
}
