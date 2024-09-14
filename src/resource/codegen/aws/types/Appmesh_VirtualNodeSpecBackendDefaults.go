package types

type Appmesh_VirtualNodeSpecBackendDefaults struct {
	// Default client policy for virtual service backends. See above for details.
	ClientPolicy Appmesh_VirtualNodeSpecBackendDefaultsClientPolicy `json:"clientPolicy,omitempty" yaml:"clientPolicy,omitempty"`
}
