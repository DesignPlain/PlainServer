package types

type Costexplorer_getCostCategoryRuleRuleNot struct {
	// Configuration block for the filter that's based on `CostCategory` values. See below.
	CostCategories []Costexplorer_getCostCategoryRuleRuleNotCostCategory `json:"costCategories,omitempty" yaml:"costCategories,omitempty"`

	// Configuration block for the specific `Dimension` to use for `Expression`. See below.
	Dimensions []Costexplorer_getCostCategoryRuleRuleNotDimension `json:"dimensions,omitempty" yaml:"dimensions,omitempty"`

	// Configuration block for the specific `Tag` to use for `Expression`. See below.
	Tags []Costexplorer_getCostCategoryRuleRuleNotTag `json:"tags,omitempty" yaml:"tags,omitempty"`
}
