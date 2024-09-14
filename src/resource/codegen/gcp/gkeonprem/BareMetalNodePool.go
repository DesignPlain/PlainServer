package gkeonprem

import types "libds/gcp/types"

type BareMetalNodePool struct {
	// The cluster this node pool belongs to.
	BareMetalCluster string `json:"bareMetalCluster,omitempty" yaml:"bareMetalCluster,omitempty"`

	// The display name for the Bare Metal Node Pool.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// The location of the resource.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The bare metal node pool name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Node pool configuration.
	   Structure is documented below.
	*/
	NodePoolConfig types.Gkeonprem_BareMetalNodePoolNodePoolConfig `json:"nodePoolConfig,omitempty" yaml:"nodePoolConfig,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Annotations on the Bare Metal Node Pool.
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
}
