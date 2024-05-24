package types

type Costexplorer_CostCategoryRuleInheritedValue struct {
	// Key to extract cost category values.
	DimensionKey string `json:"dimensionKey,omitempty" yaml:"dimensionKey,omitempty"`

	// Name of the dimension that's used to group costs. If you specify `LINKED_ACCOUNT_NAME`, the cost category value is based on account name. If you specify `TAG`, the cost category value will be based on the value of the specified tag key. Valid values are `LINKED_ACCOUNT_NAME`, `TAG`
	DimensionName string `json:"dimensionName,omitempty" yaml:"dimensionName,omitempty"`
}
