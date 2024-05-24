package types

type Vpclattice_ListenerRuleMatchHttpMatchPathMatch struct {
	// Indicates whether the match is case sensitive. Defaults to false.
	CaseSensitive bool `json:"caseSensitive,omitempty" yaml:"caseSensitive,omitempty"`

	// The header match type.
	Match Vpclattice_ListenerRuleMatchHttpMatchPathMatchMatch `json:"match,omitempty" yaml:"match,omitempty"`
}
