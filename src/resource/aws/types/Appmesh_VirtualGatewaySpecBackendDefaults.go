package types

type Appmesh_VirtualGatewaySpecBackendDefaults struct {
	// Default client policy for virtual gateway backends.
	ClientPolicy Appmesh_VirtualGatewaySpecBackendDefaultsClientPolicy `json:"clientPolicy,omitempty" yaml:"clientPolicy,omitempty"`
}
