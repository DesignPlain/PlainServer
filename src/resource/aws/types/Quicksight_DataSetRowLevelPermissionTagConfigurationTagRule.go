package types

type Quicksight_DataSetRowLevelPermissionTagConfigurationTagRule struct {
	// Column name that a tag key is assigned to.
	ColumnName string `json:"columnName,omitempty" yaml:"columnName,omitempty"`

	// A string that you want to use to filter by all the values in a column in the dataset and don’t want to list the values one by one.
	MatchAllValue string `json:"matchAllValue,omitempty" yaml:"matchAllValue,omitempty"`

	// Unique key for a tag.
	TagKey string `json:"tagKey,omitempty" yaml:"tagKey,omitempty"`

	// A string that you want to use to delimit the values when you pass the values at run time.
	TagMultiValueDelimiter string `json:"tagMultiValueDelimiter,omitempty" yaml:"tagMultiValueDelimiter,omitempty"`
}
