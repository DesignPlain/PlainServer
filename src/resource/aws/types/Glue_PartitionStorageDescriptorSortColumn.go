package types

type Glue_PartitionStorageDescriptorSortColumn struct {
	// The name of the column.
	Column string `json:"column,omitempty" yaml:"column,omitempty"`

	// Indicates that the column is sorted in ascending order (== 1), or in descending order (==0).
	SortOrder int `json:"sortOrder,omitempty" yaml:"sortOrder,omitempty"`
}
