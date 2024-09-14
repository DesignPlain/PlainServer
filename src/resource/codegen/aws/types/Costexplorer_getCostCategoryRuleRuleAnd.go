package types

type Costexplorer_getCostCategoryRuleRuleAnd struct {
	// Return results that match both `Dimension` objects.
	Ands []Costexplorer_getCostCategoryRuleRuleAndAnd `json:"ands,omitempty" yaml:"ands,omitempty"`

	// Configuration block for the filter that's based on `CostCategory` values. See below.
	CostCategories []Costexplorer_getCostCategoryRuleRuleAndCostCategory `json:"costCategories,omitempty" yaml:"costCategories,omitempty"`

	// Configuration block for the specific `Dimension` to use for `Expression`. See below.
	Dimensions []Costexplorer_getCostCategoryRuleRuleAndDimension `json:"dimensions,omitempty" yaml:"dimensions,omitempty"`

	// Return results that do not match the `Dimension` object.
	Nots []Costexplorer_getCostCategoryRuleRuleAndNot `json:"nots,omitempty" yaml:"nots,omitempty"`

	// Return results that match either `Dimension` object.
	Ors []Costexplorer_getCostCategoryRuleRuleAndOr `json:"ors,omitempty" yaml:"ors,omitempty"`

	// Configuration block for the specific `Tag` to use for `Expression`. See below.
	Tags []Costexplorer_getCostCategoryRuleRuleAndTag `json:"tags,omitempty" yaml:"tags,omitempty"`
}
