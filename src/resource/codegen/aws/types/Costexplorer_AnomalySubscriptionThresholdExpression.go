package types

type Costexplorer_AnomalySubscriptionThresholdExpression struct {
	// Return results that match both Dimension object.
	Not Costexplorer_AnomalySubscriptionThresholdExpressionNot `json:"not,omitempty" yaml:"not,omitempty"`

	// Return results that match both Dimension object.
	Ors []Costexplorer_AnomalySubscriptionThresholdExpressionOr `json:"ors,omitempty" yaml:"ors,omitempty"`

	// Configuration block for the specific Tag to use for. See Tags below.
	Tags Costexplorer_AnomalySubscriptionThresholdExpressionTags `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Return results that match both Dimension objects.
	Ands []Costexplorer_AnomalySubscriptionThresholdExpressionAnd `json:"ands,omitempty" yaml:"ands,omitempty"`

	// Configuration block for the filter that's based on  values. See Cost Category below.
	CostCategory Costexplorer_AnomalySubscriptionThresholdExpressionCostCategory `json:"costCategory,omitempty" yaml:"costCategory,omitempty"`

	// Configuration block for the specific Dimension to use for.
	Dimension Costexplorer_AnomalySubscriptionThresholdExpressionDimension `json:"dimension,omitempty" yaml:"dimension,omitempty"`
}
