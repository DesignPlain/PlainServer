package types

type Costexplorer_AnomalySubscriptionThresholdExpressionNotCostCategory struct {
	// Match options that you can use to filter your results. MatchOptions is only applicable for actions related to cost category. The default values for MatchOptions is `EQUALS` and `CASE_SENSITIVE`. Valid values are: `EQUALS`,  `ABSENT`, `STARTS_WITH`, `ENDS_WITH`, `CONTAINS`, `CASE_SENSITIVE`, `CASE_INSENSITIVE`.
	MatchOptions []string `json:"matchOptions,omitempty" yaml:"matchOptions,omitempty"`

	// Specific value of the Cost Category.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	// Unique name of the Cost Category.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
