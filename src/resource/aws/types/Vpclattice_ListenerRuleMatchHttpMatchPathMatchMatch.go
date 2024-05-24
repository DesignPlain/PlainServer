package types

type Vpclattice_ListenerRuleMatchHttpMatchPathMatchMatch struct {
	// Specifies a prefix type match. Matches the value with the prefix.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Specifies an exact type match.
	Exact string `json:"exact,omitempty" yaml:"exact,omitempty"`
}
