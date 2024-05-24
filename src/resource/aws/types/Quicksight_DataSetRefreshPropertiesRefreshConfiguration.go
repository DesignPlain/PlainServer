package types

type Quicksight_DataSetRefreshPropertiesRefreshConfiguration struct {
	// The incremental refresh for the data set. See incremental_refresh.
	IncrementalRefresh Quicksight_DataSetRefreshPropertiesRefreshConfigurationIncrementalRefresh `json:"incrementalRefresh,omitempty" yaml:"incrementalRefresh,omitempty"`
}
