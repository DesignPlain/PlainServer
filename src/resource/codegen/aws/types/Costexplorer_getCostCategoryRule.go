package types

type Costexplorer_getCostCategoryRule struct {
	// Configuration block for the value the line item is categorized as if the line item contains the matched dimension. See below.
	InheritedValues []Costexplorer_getCostCategoryRuleInheritedValue `json:"inheritedValues,omitempty" yaml:"inheritedValues,omitempty"`

	// Configuration block for the `Expression` object used to categorize costs. See below.
	Rules []Costexplorer_getCostCategoryRuleRule `json:"rules,omitempty" yaml:"rules,omitempty"`

	// Parameter type.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Default value for the cost category.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
