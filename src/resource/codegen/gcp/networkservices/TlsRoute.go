package networkservices

import types "libds/gcp/types"

type TlsRoute struct {
	// Name of the TlsRoute resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Rules that define how traffic is routed and handled.
	   Structure is documented below.
	*/
	Rules []types.Networkservices_TlsRouteRule `json:"rules,omitempty" yaml:"rules,omitempty"`

	// A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests served by the gateway.
	   Each gateway reference should match the pattern: projects/-/locations/global/gateways/<gateway_name>
	*/
	Gateways []string `json:"gateways,omitempty" yaml:"gateways,omitempty"`

	/*
	   Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served by the mesh.
	   Each mesh reference should match the pattern: projects/-/locations/global/meshes/<mesh_name>
	   The attached Mesh should be of a type SIDECAR
	*/
	Meshes []string `json:"meshes,omitempty" yaml:"meshes,omitempty"`
}
