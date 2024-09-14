package lb

import types "libds/aws/types"

type ListenerRule struct {
	// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
	Conditions []types.Lb_ListenerRuleCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`

	// The ARN of the listener to which to attach the rule.
	ListenerArn string `json:"listenerArn,omitempty" yaml:"listenerArn,omitempty"`

	// The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// An Action block. Action blocks are documented below.
	Actions []types.Lb_ListenerRuleAction `json:"actions,omitempty" yaml:"actions,omitempty"`
}
