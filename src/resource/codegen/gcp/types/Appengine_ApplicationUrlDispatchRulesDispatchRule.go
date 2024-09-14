package types

type Appengine_ApplicationUrlDispatchRulesDispatchRule struct {
	/*
	   Pathname within the host. Must start with a "/". A single "-" can be included at the end of the path.
	   The sum of the lengths of the domain and path may not exceed 100 characters.

	   - - -
	*/
	Service string `json:"service,omitempty" yaml:"service,omitempty"`

	/*
	   Domain name to match against. The wildcard "-" is supported if specified before a period: "-.".
	   Defaults to matching all domains: "-".
	*/
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`

	/*
	   Pathname within the host. Must start with a "/". A single "-" can be included at the end of the path.
	   The sum of the lengths of the domain and path may not exceed 100 characters.
	*/
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}
