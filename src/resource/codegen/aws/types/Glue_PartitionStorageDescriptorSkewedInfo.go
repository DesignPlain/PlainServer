package types

type Glue_PartitionStorageDescriptorSkewedInfo struct {
	// A list of names of columns that contain skewed values.
	SkewedColumnNames []string `json:"skewedColumnNames,omitempty" yaml:"skewedColumnNames,omitempty"`

	// A list of values that appear so frequently as to be considered skewed.
	SkewedColumnValueLocationMaps map[string]string `json:"skewedColumnValueLocationMaps,omitempty" yaml:"skewedColumnValueLocationMaps,omitempty"`

	// A map of skewed values to the columns that contain them.
	SkewedColumnValues []string `json:"skewedColumnValues,omitempty" yaml:"skewedColumnValues,omitempty"`
}
