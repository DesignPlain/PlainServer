package types

type Costexplorer_CostCategoryRuleRule struct {
	// Return results that match both `Dimension` objects.
	Ands []Costexplorer_CostCategoryRuleRuleAnd `json:"ands,omitempty" yaml:"ands,omitempty"`

	// Configuration block for the filter that's based on `CostCategory` values. See below.
	CostCategory Costexplorer_CostCategoryRuleRuleCostCategory `json:"costCategory,omitempty" yaml:"costCategory,omitempty"`

	// Configuration block for the specific `Dimension` to use for `Expression`. See below.
	Dimension Costexplorer_CostCategoryRuleRuleDimension `json:"dimension,omitempty" yaml:"dimension,omitempty"`

	// Return results that match both `Dimension` object.
	Not Costexplorer_CostCategoryRuleRuleNot `json:"not,omitempty" yaml:"not,omitempty"`

	// Return results that match both `Dimension` object.
	Ors []Costexplorer_CostCategoryRuleRuleOr `json:"ors,omitempty" yaml:"ors,omitempty"`

	// Configuration block for the specific `Tag` to use for `Expression`. See below.
	Tags Costexplorer_CostCategoryRuleRuleTags `json:"tags,omitempty" yaml:"tags,omitempty"`
}
