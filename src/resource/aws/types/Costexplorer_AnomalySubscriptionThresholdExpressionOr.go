package types

type Costexplorer_AnomalySubscriptionThresholdExpressionOr struct {
	// Configuration block for the specific Dimension to use for.
	Dimension Costexplorer_AnomalySubscriptionThresholdExpressionOrDimension `json:"dimension,omitempty" yaml:"dimension,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags Costexplorer_AnomalySubscriptionThresholdExpressionOrTags `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configuration block for the filter that's based on  values. See Cost Category below.
	CostCategory Costexplorer_AnomalySubscriptionThresholdExpressionOrCostCategory `json:"costCategory,omitempty" yaml:"costCategory,omitempty"`
}
