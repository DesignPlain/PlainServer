package types

type Appmesh_getRouteSpecGrpcRouteMatchMetadataMatch struct {
	//
	Suffix string `json:"suffix,omitempty" yaml:"suffix,omitempty"`

	//
	Exact string `json:"exact,omitempty" yaml:"exact,omitempty"`

	//
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	//
	Ranges []Appmesh_getRouteSpecGrpcRouteMatchMetadataMatchRange `json:"ranges,omitempty" yaml:"ranges,omitempty"`

	//
	Regex string `json:"regex,omitempty" yaml:"regex,omitempty"`
}
