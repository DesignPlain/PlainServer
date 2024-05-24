package types

type Costexplorer_getCostCategoryRuleRuleCostCategory struct {
	// Parameter values.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	// Key for the tag.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Match options that you can use to filter your results. MatchOptions is only applicable for actions related to cost category. The default values for MatchOptions is `EQUALS` and `CASE_SENSITIVE`. Valid values are: `EQUALS`,  `ABSENT`, `STARTS_WITH`, `ENDS_WITH`, `CONTAINS`, `CASE_SENSITIVE`, `CASE_INSENSITIVE`.
	MatchOptions []string `json:"matchOptions,omitempty" yaml:"matchOptions,omitempty"`
}
