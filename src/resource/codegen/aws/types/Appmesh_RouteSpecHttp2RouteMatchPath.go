package types

type Appmesh_RouteSpecHttp2RouteMatchPath struct {
	// The exact path to match on.
	Exact string `json:"exact,omitempty" yaml:"exact,omitempty"`

	// The regex used to match the path.
	Regex string `json:"regex,omitempty" yaml:"regex,omitempty"`
}
