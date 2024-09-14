package types

type Costexplorer_getCostCategoryRuleRuleOr struct {
	// Return results that do not match the `Dimension` object.
	Nots []Costexplorer_getCostCategoryRuleRuleOrNot `json:"nots,omitempty" yaml:"nots,omitempty"`

	// Return results that match either `Dimension` object.
	Ors []Costexplorer_getCostCategoryRuleRuleOrOr `json:"ors,omitempty" yaml:"ors,omitempty"`

	// Configuration block for the specific `Tag` to use for `Expression`. See below.
	Tags []Costexplorer_getCostCategoryRuleRuleOrTag `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Return results that match both `Dimension` objects.
	Ands []Costexplorer_getCostCategoryRuleRuleOrAnd `json:"ands,omitempty" yaml:"ands,omitempty"`

	// Configuration block for the filter that's based on `CostCategory` values. See below.
	CostCategories []Costexplorer_getCostCategoryRuleRuleOrCostCategory `json:"costCategories,omitempty" yaml:"costCategories,omitempty"`

	// Configuration block for the specific `Dimension` to use for `Expression`. See below.
	Dimensions []Costexplorer_getCostCategoryRuleRuleOrDimension `json:"dimensions,omitempty" yaml:"dimensions,omitempty"`
}
