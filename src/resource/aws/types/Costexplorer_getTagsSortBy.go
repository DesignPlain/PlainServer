package types

type Costexplorer_getTagsSortBy struct {
	// key that's used to sort the data. Valid values are: `BlendedCost`,  `UnblendedCost`, `AmortizedCost`, `NetAmortizedCost`, `NetUnblendedCost`, `UsageQuantity`, `NormalizedUsageAmount`.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// order that's used to sort the data. Valid values are: `ASCENDING`,  `DESCENDING`.
	SortOrder string `json:"sortOrder,omitempty" yaml:"sortOrder,omitempty"`
}
