package types

type Costexplorer_CostCategoryRuleRuleOr struct {
	// Configuration block for the specific `Dimension` to use for `Expression`. See below.
	Dimension Costexplorer_CostCategoryRuleRuleOrDimension `json:"dimension,omitempty" yaml:"dimension,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags Costexplorer_CostCategoryRuleRuleOrTags `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configuration block for the filter that's based on `CostCategory` values. See below.
	CostCategory Costexplorer_CostCategoryRuleRuleOrCostCategory `json:"costCategory,omitempty" yaml:"costCategory,omitempty"`
}
