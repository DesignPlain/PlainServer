package networkservices

import types "libds/gcp/types"

type TcpRoute struct {
	// Name of the TcpRoute resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Rules that define how traffic is routed and handled. At least one RouteRule must be supplied.
	   If there are multiple rules then the action taken will be the first rule to match.
	   Structure is documented below.
	*/
	Rules []types.Networkservices_TcpRouteRule `json:"rules,omitempty" yaml:"rules,omitempty"`

	// A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Gateways defines a list of gateways this TcpRoute is attached to, as one of the routing rules to route the requests served by the gateway.
	   Each gateway reference should match the pattern: projects/-/locations/global/gateways/<gateway_name>
	*/
	Gateways []string `json:"gateways,omitempty" yaml:"gateways,omitempty"`

	/*
	   Set of label tags associated with the TcpRoute resource.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Meshes defines a list of meshes this TcpRoute is attached to, as one of the routing rules to route the requests served by the mesh.
	   Each mesh reference should match the pattern: projects/-/locations/global/meshes/<mesh_name>
	   The attached Mesh should be of a type SIDECAR
	*/
	Meshes []string `json:"meshes,omitempty" yaml:"meshes,omitempty"`
}
