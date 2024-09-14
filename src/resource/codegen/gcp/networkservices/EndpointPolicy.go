package networkservices

import types "libds/gcp/types"

type EndpointPolicy struct {
	/*
	   Set of label tags associated with the TcpRoute resource.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Port selector for the (matched) endpoints. If no port selector is provided, the matched config is applied to all ports.
	   Structure is documented below.
	*/
	TrafficPortSelector types.Networkservices_EndpointPolicyTrafficPortSelector `json:"trafficPortSelector,omitempty" yaml:"trafficPortSelector,omitempty"`

	/*
	   The type of endpoint policy. This is primarily used to validate the configuration.
	   Possible values are: `SIDECAR_PROXY`, `GRPC_SERVER`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// This field specifies the URL of AuthorizationPolicy resource that applies authorization policies to the inbound traffic at the matched endpoints.
	AuthorizationPolicy string `json:"authorizationPolicy,omitempty" yaml:"authorizationPolicy,omitempty"`

	// A URL referring to a ClientTlsPolicy resource. ClientTlsPolicy can be set to specify the authentication for traffic from the proxy to the actual endpoints.
	ClientTlsPolicy string `json:"clientTlsPolicy,omitempty" yaml:"clientTlsPolicy,omitempty"`

	// A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Required. A matcher that selects endpoints to which the policies should be applied.
	   Structure is documented below.
	*/
	EndpointMatcher types.Networkservices_EndpointPolicyEndpointMatcher `json:"endpointMatcher,omitempty" yaml:"endpointMatcher,omitempty"`

	// Name of the EndpointPolicy resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A URL referring to ServerTlsPolicy resource. ServerTlsPolicy is used to determine the authentication policy to be applied to terminate the inbound traffic at the identified backends.
	ServerTlsPolicy string `json:"serverTlsPolicy,omitempty" yaml:"serverTlsPolicy,omitempty"`
}
