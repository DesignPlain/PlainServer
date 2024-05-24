package types

type Appmesh_getRouteSpecHttpRouteMatchHeaderMatch struct {
	//
	Exact string `json:"exact,omitempty" yaml:"exact,omitempty"`

	//
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	//
	Ranges []Appmesh_getRouteSpecHttpRouteMatchHeaderMatchRange `json:"ranges,omitempty" yaml:"ranges,omitempty"`

	//
	Regex string `json:"regex,omitempty" yaml:"regex,omitempty"`

	//
	Suffix string `json:"suffix,omitempty" yaml:"suffix,omitempty"`
}
