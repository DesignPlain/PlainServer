package types

type Appmesh_RouteSpecHttpRouteMatchPath struct {
	// The exact path to match on.
	Exact string `json:"exact,omitempty" yaml:"exact,omitempty"`

	// The regex used to match the path.
	Regex string `json:"regex,omitempty" yaml:"regex,omitempty"`
}
