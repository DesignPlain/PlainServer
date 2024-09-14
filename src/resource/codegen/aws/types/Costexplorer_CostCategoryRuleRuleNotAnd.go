package types

type Costexplorer_CostCategoryRuleRuleNotAnd struct {
	// Configuration block for the specific `Dimension` to use for `Expression`. See below.
	Dimension Costexplorer_CostCategoryRuleRuleNotAndDimension `json:"dimension,omitempty" yaml:"dimension,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags Costexplorer_CostCategoryRuleRuleNotAndTags `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configuration block for the filter that's based on `CostCategory` values. See below.
	CostCategory Costexplorer_CostCategoryRuleRuleNotAndCostCategory `json:"costCategory,omitempty" yaml:"costCategory,omitempty"`
}
