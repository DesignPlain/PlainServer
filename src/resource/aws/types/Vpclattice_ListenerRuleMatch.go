package types

type Vpclattice_ListenerRuleMatch struct {
	// The HTTP criteria that a rule must match.
	HttpMatch Vpclattice_ListenerRuleMatchHttpMatch `json:"httpMatch,omitempty" yaml:"httpMatch,omitempty"`
}
