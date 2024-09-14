package types

type Glue_CatalogTableStorageDescriptorSkewedInfo struct {
	// List of names of columns that contain skewed values.
	SkewedColumnNames []string `json:"skewedColumnNames,omitempty" yaml:"skewedColumnNames,omitempty"`

	// List of values that appear so frequently as to be considered skewed.
	SkewedColumnValueLocationMaps map[string]string `json:"skewedColumnValueLocationMaps,omitempty" yaml:"skewedColumnValueLocationMaps,omitempty"`

	// Map of skewed values to the columns that contain them.
	SkewedColumnValues []string `json:"skewedColumnValues,omitempty" yaml:"skewedColumnValues,omitempty"`
}
