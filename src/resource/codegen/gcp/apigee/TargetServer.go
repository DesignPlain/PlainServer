package apigee

import types "libds/gcp/types"

type TargetServer struct {
	/*
	   The Apigee environment group associated with the Apigee environment,
	   in the format `organizations/{{org_name}}/environments/{{env_name}}`.


	   - - -
	*/
	EnvId string `json:"envId,omitempty" yaml:"envId,omitempty"`

	// The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	// Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or more TargetServers need to taken out of rotation periodically. Defaults to true.
	IsEnabled bool `json:"isEnabled,omitempty" yaml:"isEnabled,omitempty"`

	// The resource id of this reference. Values must match the regular expression [\w\s-.]+.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	/*
	   Immutable. The protocol used by this TargetServer.
	   Possible values are: `HTTP`, `HTTP2`, `GRPC_TARGET`, `GRPC`, `EXTERNAL_CALLOUT`.
	*/
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	/*
	   Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
	   Structure is documented below.
	*/
	SSlInfo types.Apigee_TargetServerSSlInfo `json:"sSlInfo,omitempty" yaml:"sSlInfo,omitempty"`

	// A human-readable description of this TargetServer.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
