package types

type Costexplorer_CostCategoryRule struct {
	// You can define the CostCategoryRule rule type as either `REGULAR` or `INHERITED_VALUE`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Default value for the cost category.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// Configuration block for the value the line item is categorized as if the line item contains the matched dimension. See below.
	InheritedValue Costexplorer_CostCategoryRuleInheritedValue `json:"inheritedValue,omitempty" yaml:"inheritedValue,omitempty"`

	// Configuration block for the `Expression` object used to categorize costs. See below.
	Rule Costexplorer_CostCategoryRuleRule `json:"rule,omitempty" yaml:"rule,omitempty"`
}
