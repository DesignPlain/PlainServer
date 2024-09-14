package types

type Ecr_getLifecyclePolicyDocumentRule struct {
	// Specifies the action type.
	Action Ecr_getLifecyclePolicyDocumentRuleAction `json:"action,omitempty" yaml:"action,omitempty"`

	// Describes the purpose of a rule within a lifecycle policy.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Sets the order in which rules are evaluated, lowest to highest. When you add rules to a lifecycle policy, you must give them each a unique value for `priority`. Values do not need to be sequential across rules in a policy. A rule with a `tag_status` value of "any" must have the highest value for `priority` and be evaluated last.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// Collects parameters describing the selection criteria for the ECR lifecycle policy:
	Selection Ecr_getLifecyclePolicyDocumentRuleSelection `json:"selection,omitempty" yaml:"selection,omitempty"`
}
