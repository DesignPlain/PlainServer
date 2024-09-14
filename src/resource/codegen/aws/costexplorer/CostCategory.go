package costexplorer

import types "libds/aws/types"

type CostCategory struct {
	/*
	   The Cost Category's effective start date. It can only be a billing start date (first day of the month). If the date isn't provided, it's the first day of the current month. Dates can't be before the previous twelve months, or in the future. For example `2022-11-01T00:00:00Z`.

	   The following arguments are optional:
	*/
	EffectiveStart string `json:"effectiveStart,omitempty" yaml:"effectiveStart,omitempty"`

	// Unique name for the Cost Category.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Rule schema version in this particular Cost Category.
	RuleVersion string `json:"ruleVersion,omitempty" yaml:"ruleVersion,omitempty"`

	// Configuration block for the Cost Category rules used to categorize costs. See below.
	Rules []types.Costexplorer_CostCategoryRule `json:"rules,omitempty" yaml:"rules,omitempty"`

	// Configuration block for the split charge rules used to allocate your charges between your Cost Category values. See below.
	SplitChargeRules []types.Costexplorer_CostCategorySplitChargeRule `json:"splitChargeRules,omitempty" yaml:"splitChargeRules,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Default value for the cost category.
	DefaultValue string `json:"defaultValue,omitempty" yaml:"defaultValue,omitempty"`
}
