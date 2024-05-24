package types

type Quicksight_DataSetRefreshPropertiesRefreshConfigurationIncrementalRefresh struct {
	// The lookback window setup for an incremental refresh configuration. See lookback_window.
	LookbackWindow Quicksight_DataSetRefreshPropertiesRefreshConfigurationIncrementalRefreshLookbackWindow `json:"lookbackWindow,omitempty" yaml:"lookbackWindow,omitempty"`
}
