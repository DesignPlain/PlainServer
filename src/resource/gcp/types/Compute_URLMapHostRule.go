package types

type Compute_URLMapHostRule struct {
	/*
	   An optional description of this resource. Provide this property when you create
	   the resource.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The list of host patterns to match. They must be valid hostnames, except - will
	   match any string of ([a-z0-9-.]-). In that case, - must be the first character
	   and must be followed in the pattern by either - or ..
	*/
	Hosts []string `json:"hosts,omitempty" yaml:"hosts,omitempty"`

	/*
	   The name of the PathMatcher to use to match the path portion of the URL if the
	   hostRule matches the URL's host portion.
	*/
	PathMatcher string `json:"pathMatcher,omitempty" yaml:"pathMatcher,omitempty"`
}
