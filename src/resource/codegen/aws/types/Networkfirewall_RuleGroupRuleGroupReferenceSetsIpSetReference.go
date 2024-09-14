package types

type Networkfirewall_RuleGroupRuleGroupReferenceSetsIpSetReference struct {
	// Set of configuration blocks that define the IP Reference information. See IP Set Reference below for details.
	IpSetReferences []Networkfirewall_RuleGroupRuleGroupReferenceSetsIpSetReferenceIpSetReference `json:"ipSetReferences,omitempty" yaml:"ipSetReferences,omitempty"`

	//
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
