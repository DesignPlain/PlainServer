package types

type Appmesh_RouteSpecHttpRouteMatchHeaderMatch struct {
	// Header value sent by the client must begin with the specified characters.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Object that specifies the range of numbers that the header value sent by the client must be included in.
	Range Appmesh_RouteSpecHttpRouteMatchHeaderMatchRange `json:"range,omitempty" yaml:"range,omitempty"`

	// The regex used to match the path.
	Regex string `json:"regex,omitempty" yaml:"regex,omitempty"`

	// Header value sent by the client must end with the specified characters.
	Suffix string `json:"suffix,omitempty" yaml:"suffix,omitempty"`

	// The exact path to match on.
	Exact string `json:"exact,omitempty" yaml:"exact,omitempty"`
}
