package types

type Quicksight_DataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindow struct {
	// The name of the lookback window column.
	ColumnName string `json:"columnName,omitempty" yaml:"columnName,omitempty"`

	// The lookback window column size.
	Size int `json:"size,omitempty" yaml:"size,omitempty"`

	// The size unit that is used for the lookback window column. Valid values for this structure are `HOUR`, `DAY`, and `WEEK`.
	SizeUnit string `json:"sizeUnit,omitempty" yaml:"sizeUnit,omitempty"`
}
