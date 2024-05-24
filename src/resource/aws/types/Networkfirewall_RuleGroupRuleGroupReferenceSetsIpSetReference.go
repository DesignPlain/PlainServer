package types

type Networkfirewall_RuleGroupRuleGroupReferenceSetsIpSetReference struct {
	// A unique alphanumeric string to identify the `ip_set`.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Set of configuration blocks that define the IP Reference information. See IP Set Reference below for details.
	IpSetReferences []Networkfirewall_RuleGroupRuleGroupReferenceSetsIpSetReferenceIpSetReference `json:"ipSetReferences,omitempty" yaml:"ipSetReferences,omitempty"`
}
