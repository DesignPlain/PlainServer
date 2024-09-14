package types

type Appmesh_RouteSpecGrpcRouteMatchMetadataMatch struct {
	// Value sent by the client must begin with the specified characters. Must be between 1 and 255 characters in length.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Object that specifies the range of numbers that the value sent by the client must be included in.
	Range Appmesh_RouteSpecGrpcRouteMatchMetadataMatchRange `json:"range,omitempty" yaml:"range,omitempty"`

	// Value sent by the client must include the specified characters. Must be between 1 and 255 characters in length.
	Regex string `json:"regex,omitempty" yaml:"regex,omitempty"`

	// Value sent by the client must end with the specified characters. Must be between 1 and 255 characters in length.
	Suffix string `json:"suffix,omitempty" yaml:"suffix,omitempty"`

	// Value sent by the client must match the specified value exactly. Must be between 1 and 255 characters in length.
	Exact string `json:"exact,omitempty" yaml:"exact,omitempty"`
}
