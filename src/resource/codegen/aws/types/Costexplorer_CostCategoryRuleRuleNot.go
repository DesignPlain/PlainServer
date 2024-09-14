package types

type Costexplorer_CostCategoryRuleRuleNot struct {
	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags Costexplorer_CostCategoryRuleRuleNotTags `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Return results that match both `Dimension` objects.
	Ands []Costexplorer_CostCategoryRuleRuleNotAnd `json:"ands,omitempty" yaml:"ands,omitempty"`

	// Configuration block for the filter that's based on `CostCategory` values. See below.
	CostCategory Costexplorer_CostCategoryRuleRuleNotCostCategory `json:"costCategory,omitempty" yaml:"costCategory,omitempty"`

	// Configuration block for the specific `Dimension` to use for `Expression`. See below.
	Dimension Costexplorer_CostCategoryRuleRuleNotDimension `json:"dimension,omitempty" yaml:"dimension,omitempty"`

	// Return results that match both `Dimension` object.
	Not Costexplorer_CostCategoryRuleRuleNotNot `json:"not,omitempty" yaml:"not,omitempty"`

	// Return results that match both `Dimension` object.
	Ors []Costexplorer_CostCategoryRuleRuleNotOr `json:"ors,omitempty" yaml:"ors,omitempty"`
}
