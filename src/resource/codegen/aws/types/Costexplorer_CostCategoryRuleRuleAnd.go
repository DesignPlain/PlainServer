package types

type Costexplorer_CostCategoryRuleRuleAnd struct {
	// Configuration block for the specific `Dimension` to use for `Expression`. See below.
	Dimension Costexplorer_CostCategoryRuleRuleAndDimension `json:"dimension,omitempty" yaml:"dimension,omitempty"`

	// Return results that match both `Dimension` object.
	Not Costexplorer_CostCategoryRuleRuleAndNot `json:"not,omitempty" yaml:"not,omitempty"`

	// Return results that match both `Dimension` object.
	Ors []Costexplorer_CostCategoryRuleRuleAndOr `json:"ors,omitempty" yaml:"ors,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags Costexplorer_CostCategoryRuleRuleAndTags `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Return results that match both `Dimension` objects.
	Ands []Costexplorer_CostCategoryRuleRuleAndAnd `json:"ands,omitempty" yaml:"ands,omitempty"`

	// Configuration block for the filter that's based on `CostCategory` values. See below.
	CostCategory Costexplorer_CostCategoryRuleRuleAndCostCategory `json:"costCategory,omitempty" yaml:"costCategory,omitempty"`
}
