package types

type Costexplorer_getCostCategorySplitChargeRule struct {
	// Method that's used to define how to split your source costs across your targets. Valid values are `FIXED`, `PROPORTIONAL`, `EVEN`
	Method string `json:"method,omitempty" yaml:"method,omitempty"`

	// Configuration block for the parameters for a split charge method. This is only required for the `FIXED` method. See below.
	Parameters []Costexplorer_getCostCategorySplitChargeRuleParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Cost Category value that you want to split.
	Source string `json:"source,omitempty" yaml:"source,omitempty"`

	// Cost Category values that you want to split costs across. These values can't be used as a source in other split charge rules.
	Targets []string `json:"targets,omitempty" yaml:"targets,omitempty"`
}
