package types

type Vpclattice_ListenerRuleMatchHttpMatchHeaderMatch struct {
	// Indicates whether the match is case sensitive. Defaults to false.
	CaseSensitive bool `json:"caseSensitive,omitempty" yaml:"caseSensitive,omitempty"`

	// The header match type.
	Match Vpclattice_ListenerRuleMatchHttpMatchHeaderMatchMatch `json:"match,omitempty" yaml:"match,omitempty"`

	// The name of the header.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
