package types

type Appmesh_getRouteSpecGrpcRouteMatchMetadata struct {
	//
	Invert bool `json:"invert,omitempty" yaml:"invert,omitempty"`

	//
	Matches []Appmesh_getRouteSpecGrpcRouteMatchMetadataMatch `json:"matches,omitempty" yaml:"matches,omitempty"`

	// Name of the route.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
