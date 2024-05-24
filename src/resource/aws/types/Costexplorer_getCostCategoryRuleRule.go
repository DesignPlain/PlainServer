package types

type Costexplorer_getCostCategoryRuleRule struct {
	// Return results that match both `Dimension` objects.
	Ands []Costexplorer_getCostCategoryRuleRuleAnd `json:"ands,omitempty" yaml:"ands,omitempty"`

	// Configuration block for the filter that's based on `CostCategory` values. See below.
	CostCategories []Costexplorer_getCostCategoryRuleRuleCostCategory `json:"costCategories,omitempty" yaml:"costCategories,omitempty"`

	// Configuration block for the specific `Dimension` to use for `Expression`. See below.
	Dimensions []Costexplorer_getCostCategoryRuleRuleDimension `json:"dimensions,omitempty" yaml:"dimensions,omitempty"`

	// Return results that do not match the `Dimension` object.
	Nots []Costexplorer_getCostCategoryRuleRuleNot `json:"nots,omitempty" yaml:"nots,omitempty"`

	// Return results that match either `Dimension` object.
	Ors []Costexplorer_getCostCategoryRuleRuleOr `json:"ors,omitempty" yaml:"ors,omitempty"`

	// Configuration block for the specific `Tag` to use for `Expression`. See below.
	Tags []Costexplorer_getCostCategoryRuleRuleTag `json:"tags,omitempty" yaml:"tags,omitempty"`
}
