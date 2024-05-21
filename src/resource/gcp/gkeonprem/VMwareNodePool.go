package gkeonprem

import types "DesignSphere_Server/src/resource/gcp/types"

type VMwareNodePool struct {
	// The vmware node pool name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Node Pool autoscaling config for the node pool.
	   Structure is documented below.
	*/
	NodePoolAutoscaling types.Gkeonprem_VMwareNodePoolNodePoolAutoscaling `json:"nodePoolAutoscaling,omitempty" yaml:"nodePoolAutoscaling,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The cluster this node pool belongs to.
	VmwareCluster string `json:"vmwareCluster,omitempty" yaml:"vmwareCluster,omitempty"`

	/*
	   Annotations on the node Pool.
	   This field has the same restrictions as Kubernetes annotations.
	   The total size of all keys and values combined is limited to 256k.
	   Key can have 2 segments: prefix (optional) and name (required),
	   separated by a slash (/).
	   Prefix must be a DNS subdomain.
	   Name must be 63 characters or less, begin and end with alphanumerics,
	   with dashes (-), underscores (_), dots (.), and alphanumerics between.

	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	/*
	   The node configuration of the node pool.
	   Structure is documented below.
	*/
	Config types.Gkeonprem_VMwareNodePoolConfig `json:"config,omitempty" yaml:"config,omitempty"`

	// The display name for the node pool.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// The location of the resource.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}
