package types

type Costexplorer_CostCategoryRuleRuleOr struct {
	// Return results that match both `Dimension` objects.
	Ands []Costexplorer_CostCategoryRuleRuleOrAnd `json:"ands,omitempty" yaml:"ands,omitempty"`

	// Configuration block for the filter that's based on `CostCategory` values. See below.
	CostCategory Costexplorer_CostCategoryRuleRuleOrCostCategory `json:"costCategory,omitempty" yaml:"costCategory,omitempty"`

	// Configuration block for the specific `Dimension` to use for `Expression`. See below.
	Dimension Costexplorer_CostCategoryRuleRuleOrDimension `json:"dimension,omitempty" yaml:"dimension,omitempty"`

	// Return results that match both `Dimension` object.
	Not Costexplorer_CostCategoryRuleRuleOrNot `json:"not,omitempty" yaml:"not,omitempty"`

	// Return results that match both `Dimension` object.
	Ors []Costexplorer_CostCategoryRuleRuleOrOr `json:"ors,omitempty" yaml:"ors,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags Costexplorer_CostCategoryRuleRuleOrTags `json:"tags,omitempty" yaml:"tags,omitempty"`
}
