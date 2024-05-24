package types

type Appmesh_MeshSpecEgressFilter struct {
	/*
	   Egress filter type. By default, the type is `DROP_ALL`.
	   Valid values are `ALLOW_ALL` and `DROP_ALL`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
