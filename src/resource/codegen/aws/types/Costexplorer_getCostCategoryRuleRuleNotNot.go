package types

type Costexplorer_getCostCategoryRuleRuleNotNot struct {
	// Configuration block for the filter that's based on `CostCategory` values. See below.
	CostCategories []Costexplorer_getCostCategoryRuleRuleNotNotCostCategory `json:"costCategories,omitempty" yaml:"costCategories,omitempty"`

	// Configuration block for the specific `Dimension` to use for `Expression`. See below.
	Dimensions []Costexplorer_getCostCategoryRuleRuleNotNotDimension `json:"dimensions,omitempty" yaml:"dimensions,omitempty"`

	// Configuration block for the specific `Tag` to use for `Expression`. See below.
	Tags []Costexplorer_getCostCategoryRuleRuleNotNotTag `json:"tags,omitempty" yaml:"tags,omitempty"`
}
