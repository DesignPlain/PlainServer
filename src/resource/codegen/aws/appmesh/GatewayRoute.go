package appmesh

import types "libds/aws/types"

type GatewayRoute struct {
	// Gateway route specification to apply.
	Spec types.Appmesh_GatewayRouteSpec `json:"spec,omitempty" yaml:"spec,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Name of the virtual gateway to associate the gateway route with. Must be between 1 and 255 characters in length.
	VirtualGatewayName string `json:"virtualGatewayName,omitempty" yaml:"virtualGatewayName,omitempty"`

	// Name of the service mesh in which to create the gateway route. Must be between 1 and 255 characters in length.
	MeshName string `json:"meshName,omitempty" yaml:"meshName,omitempty"`

	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	MeshOwner string `json:"meshOwner,omitempty" yaml:"meshOwner,omitempty"`

	// Name to use for the gateway route. Must be between 1 and 255 characters in length.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
