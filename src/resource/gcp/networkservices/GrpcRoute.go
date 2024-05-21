package networkservices

import types "DesignSphere_Server/src/resource/gcp/types"

type GrpcRoute struct {
	// List of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway.
	Gateways []string `json:"gateways,omitempty" yaml:"gateways,omitempty"`

	// Required. Service hostnames with an optional port for which this route describes traffic.
	Hostnames []string `json:"hostnames,omitempty" yaml:"hostnames,omitempty"`

	/*
	   Set of label tags associated with the GrpcRoute resource.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// List of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh.
	Meshes []string `json:"meshes,omitempty" yaml:"meshes,omitempty"`

	// Name of the GrpcRoute resource.
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
	Rules []types.Networkservices_GrpcRouteRule `json:"rules,omitempty" yaml:"rules,omitempty"`

	// A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
