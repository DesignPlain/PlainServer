package types

type Appmesh_GatewayRouteSpecHttpRouteMatchHeaderMatch struct {
	// Header value sent by the client must match the specified value exactly.
	Exact string `json:"exact,omitempty" yaml:"exact,omitempty"`

	// Header value sent by the client must begin with the specified characters.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Object that specifies the range of numbers that the header value sent by the client must be included in.
	Range Appmesh_GatewayRouteSpecHttpRouteMatchHeaderMatchRange `json:"range,omitempty" yaml:"range,omitempty"`

	// Header value sent by the client must include the specified characters.
	Regex string `json:"regex,omitempty" yaml:"regex,omitempty"`

	// Header value sent by the client must end with the specified characters.
	Suffix string `json:"suffix,omitempty" yaml:"suffix,omitempty"`
}
