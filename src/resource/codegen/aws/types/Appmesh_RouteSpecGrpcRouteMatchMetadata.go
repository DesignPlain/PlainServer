package types

type Appmesh_RouteSpecGrpcRouteMatchMetadata struct {
	// If `true`, the match is on the opposite of the `match` criteria. Default is `false`.
	Invert bool `json:"invert,omitempty" yaml:"invert,omitempty"`

	// Data to match from the request.
	Match Appmesh_RouteSpecGrpcRouteMatchMetadataMatch `json:"match,omitempty" yaml:"match,omitempty"`

	// Name of the route. Must be between 1 and 50 characters in length.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
