package types

type Vpclattice_ListenerRuleMatchHttpMatchHeaderMatchMatch struct {
	// Specifies a contains type match.
	Contains string `json:"contains,omitempty" yaml:"contains,omitempty"`

	// Specifies an exact type match.
	Exact string `json:"exact,omitempty" yaml:"exact,omitempty"`

	// Specifies a prefix type match. Matches the value with the prefix.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
