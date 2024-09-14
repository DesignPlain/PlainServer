package types

type Networkmanager_getCoreNetworkPolicyDocumentAttachmentPolicy struct {
	// Action to take when a condition is true. Detailed Below.
	Action Networkmanager_getCoreNetworkPolicyDocumentAttachmentPolicyAction `json:"action,omitempty" yaml:"action,omitempty"`

	// Valid values include `and` or `or`. This is a mandatory parameter only if you have more than one condition. The `condition_logic` apply to all of the conditions for a rule, which also means nested conditions of `and` or `or` are not supported. Use `or` if you want to associate the attachment with the segment by either the segment name or attachment tag value, or by the chosen conditions. Use `and` if you want to associate the attachment with the segment by either the segment name or attachment tag value and by the chosen conditions. Detailed Below.
	ConditionLogic string `json:"conditionLogic,omitempty" yaml:"conditionLogic,omitempty"`

	// A block argument. Detailed Below.
	Conditions []Networkmanager_getCoreNetworkPolicyDocumentAttachmentPolicyCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`

	// A user-defined description that further helps identify the rule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// An integer from `1` to `65535` indicating the rule's order number. Rules are processed in order from the lowest numbered rule to the highest. Rules stop processing when a rule is matched. It's important to make sure that you number your rules in the exact order that you want them processed.
	RuleNumber int `json:"ruleNumber,omitempty" yaml:"ruleNumber,omitempty"`
}
