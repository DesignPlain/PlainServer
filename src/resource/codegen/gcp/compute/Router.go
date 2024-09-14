package compute

import types "libds/gcp/types"

type Router struct {
	/*
	   BGP information specific to this router.
	   Structure is documented below.
	*/
	Bgp types.Compute_RouterBgp `json:"bgp,omitempty" yaml:"bgp,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Indicates if a router is dedicated for use with encrypted VLAN
	   attachments (interconnectAttachments).
	*/
	EncryptedInterconnectRouter bool `json:"encryptedInterconnectRouter,omitempty" yaml:"encryptedInterconnectRouter,omitempty"`

	/*
	   Name of the resource. The name must be 1-63 characters long, and
	   comply with RFC1035. Specifically, the name must be 1-63 characters
	   long and match the regular expression `a-z?`
	   which means the first character must be a lowercase letter, and all
	   following characters must be a dash, lowercase letter, or digit,
	   except the last character, which cannot be a dash.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   A reference to the network to which this router belongs.


	   - - -
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Region where the router resides.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
