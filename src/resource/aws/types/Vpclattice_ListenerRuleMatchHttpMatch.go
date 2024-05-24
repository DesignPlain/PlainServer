package types

type Vpclattice_ListenerRuleMatchHttpMatch struct {
	// The header matches. Matches incoming requests with rule based on request header value before applying rule action.
	HeaderMatches []Vpclattice_ListenerRuleMatchHttpMatchHeaderMatch `json:"headerMatches,omitempty" yaml:"headerMatches,omitempty"`

	// The HTTP method type.
	Method string `json:"method,omitempty" yaml:"method,omitempty"`

	// The path match.
	PathMatch Vpclattice_ListenerRuleMatchHttpMatchPathMatch `json:"pathMatch,omitempty" yaml:"pathMatch,omitempty"`
}
