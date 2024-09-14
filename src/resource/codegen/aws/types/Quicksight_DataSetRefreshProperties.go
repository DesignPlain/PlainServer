package types

type Quicksight_DataSetRefreshProperties struct {
	// The refresh configuration for the data set. See refresh_configuration.
	RefreshConfiguration Quicksight_DataSetRefreshPropertiesRefreshConfiguration `json:"refreshConfiguration,omitempty" yaml:"refreshConfiguration,omitempty"`
}
