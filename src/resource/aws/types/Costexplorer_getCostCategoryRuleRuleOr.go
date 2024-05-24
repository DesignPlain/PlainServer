package types

type Costexplorer_getCostCategoryRuleRuleOr struct {
	// Configuration block for the filter that's based on `CostCategory` values. See below.
	CostCategories []Costexplorer_getCostCategoryRuleRuleOrCostCategory `json:"costCategories,omitempty" yaml:"costCategories,omitempty"`

	// Configuration block for the specific `Dimension` to use for `Expression`. See below.
	Dimensions []Costexplorer_getCostCategoryRuleRuleOrDimension `json:"dimensions,omitempty" yaml:"dimensions,omitempty"`

	// Configuration block for the specific `Tag` to use for `Expression`. See below.
	Tags []Costexplorer_getCostCategoryRuleRuleOrTag `json:"tags,omitempty" yaml:"tags,omitempty"`
}
