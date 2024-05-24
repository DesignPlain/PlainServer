package types

type Dlm_LifecyclePolicyPolicyDetailsSchedule struct {
	// See the `share_rule` block. Max of 1 per schedule.
	ShareRule Dlm_LifecyclePolicyPolicyDetailsScheduleShareRule `json:"shareRule,omitempty" yaml:"shareRule,omitempty"`

	// Whether to copy all user-defined tags from the source snapshot to the cross-region snapshot copy.
	CopyTags bool `json:"copyTags,omitempty" yaml:"copyTags,omitempty"`

	// See the `cross_region_copy_rule` block. Max of 3 per schedule.
	CrossRegionCopyRules []Dlm_LifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRule `json:"crossRegionCopyRules,omitempty" yaml:"crossRegionCopyRules,omitempty"`

	// See the `fast_restore_rule` block. Max of 1 per schedule.
	FastRestoreRule Dlm_LifecyclePolicyPolicyDetailsScheduleFastRestoreRule `json:"fastRestoreRule,omitempty" yaml:"fastRestoreRule,omitempty"`

	// A descriptive name for the action.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies the retention rule for cross-Region snapshot copies. See the `retain_rule` block. Max of 1 per action.
	RetainRule Dlm_LifecyclePolicyPolicyDetailsScheduleRetainRule `json:"retainRule,omitempty" yaml:"retainRule,omitempty"`

	// See the `create_rule` block. Max of 1 per schedule.
	CreateRule Dlm_LifecyclePolicyPolicyDetailsScheduleCreateRule `json:"createRule,omitempty" yaml:"createRule,omitempty"`

	// The AMI deprecation rule for cross-Region AMI copies created by the rule. See the `deprecate_rule` block.
	DeprecateRule Dlm_LifecyclePolicyPolicyDetailsScheduleDeprecateRule `json:"deprecateRule,omitempty" yaml:"deprecateRule,omitempty"`

	// A map of tag keys and their values. DLM lifecycle policies will already tag the snapshot with the tags on the volume. This configuration adds extra tags on top of these.
	TagsToAdd map[string]string `json:"tagsToAdd,omitempty" yaml:"tagsToAdd,omitempty"`

	// A map of tag keys and variable values, where the values are determined when the policy is executed. Only `$(instance-id)` or `$(timestamp)` are valid values. Can only be used when `resource_types` is `INSTANCE`.
	VariableTags map[string]string `json:"variableTags,omitempty" yaml:"variableTags,omitempty"`
}
