package types

type Appmesh_VirtualRouterSpec struct {
	/*
	   Listeners that the virtual router is expected to receive inbound traffic from.
	   Currently only one listener is supported per virtual router.
	*/
	Listeners []Appmesh_VirtualRouterSpecListener `json:"listeners,omitempty" yaml:"listeners,omitempty"`
}
