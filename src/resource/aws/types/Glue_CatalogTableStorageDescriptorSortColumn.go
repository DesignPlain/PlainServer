package types

type Glue_CatalogTableStorageDescriptorSortColumn struct {
	// Name of the column.
	Column string `json:"column,omitempty" yaml:"column,omitempty"`

	// Whether the column is sorted in ascending (`1`) or descending order (`0`).
	SortOrder int `json:"sortOrder,omitempty" yaml:"sortOrder,omitempty"`
}
