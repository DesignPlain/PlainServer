package compute

import types "DesignSphere_Server/src/resource/gcp/types"

type NetworkEndpointList struct {
	/*
	   The network endpoint group these endpoints are part of.


	   - - -
	*/
	NetworkEndpointGroup string `json:"networkEndpointGroup,omitempty" yaml:"networkEndpointGroup,omitempty"`

	/*
	   The network endpoints to be added to the enclosing network endpoint group
	   (NEG). Each endpoint specifies an IP address and port, along with
	   additional information depending on the NEG type.
	   Structure is documented below.
	*/
	NetworkEndpoints []types.Compute_NetworkEndpointListNetworkEndpoint `json:"networkEndpoints,omitempty" yaml:"networkEndpoints,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Zone where the containing network endpoint group is located.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
