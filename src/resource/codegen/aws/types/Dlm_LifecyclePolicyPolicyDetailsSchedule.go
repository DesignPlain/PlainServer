package types

type Dlm_LifecyclePolicyPolicyDetailsSchedule struct {
	//
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	RetainRule Dlm_LifecyclePolicyPolicyDetailsScheduleRetainRule `json:"retainRule,omitempty" yaml:"retainRule,omitempty"`

	//
	CopyTags bool `json:"copyTags,omitempty" yaml:"copyTags,omitempty"`

	// See the `create_rule` block. Max of 1 per schedule.
	CreateRule Dlm_LifecyclePolicyPolicyDetailsScheduleCreateRule `json:"createRule,omitempty" yaml:"createRule,omitempty"`

	// See the `fast_restore_rule` block. Max of 1 per schedule.
	FastRestoreRule Dlm_LifecyclePolicyPolicyDetailsScheduleFastRestoreRule `json:"fastRestoreRule,omitempty" yaml:"fastRestoreRule,omitempty"`

	// See the `share_rule` block. Max of 1 per schedule.
	ShareRule Dlm_LifecyclePolicyPolicyDetailsScheduleShareRule `json:"shareRule,omitempty" yaml:"shareRule,omitempty"`

	// A map of tag keys and their values. DLM lifecycle policies will already tag the snapshot with the tags on the volume. This configuration adds extra tags on top of these.
	TagsToAdd map[string]string `json:"tagsToAdd,omitempty" yaml:"tagsToAdd,omitempty"`

	// A map of tag keys and variable values, where the values are determined when the policy is executed. Only `$(instance-id)` or `$(timestamp)` are valid values. Can only be used when `resource_types` is `INSTANCE`.
	VariableTags map[string]string `json:"variableTags,omitempty" yaml:"variableTags,omitempty"`

	// See the `cross_region_copy_rule` block. Max of 3 per schedule.
	CrossRegionCopyRules []Dlm_LifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRule `json:"crossRegionCopyRules,omitempty" yaml:"crossRegionCopyRules,omitempty"`

	//
	DeprecateRule Dlm_LifecyclePolicyPolicyDetailsScheduleDeprecateRule `json:"deprecateRule,omitempty" yaml:"deprecateRule,omitempty"`
}
