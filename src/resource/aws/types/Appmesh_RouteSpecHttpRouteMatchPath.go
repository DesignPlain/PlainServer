package types

type Appmesh_RouteSpecHttpRouteMatchPath struct {
	// Header value sent by the client must include the specified characters.
	Regex string `json:"regex,omitempty" yaml:"regex,omitempty"`

	// Header value sent by the client must match the specified value exactly.
	Exact string `json:"exact,omitempty" yaml:"exact,omitempty"`
}
