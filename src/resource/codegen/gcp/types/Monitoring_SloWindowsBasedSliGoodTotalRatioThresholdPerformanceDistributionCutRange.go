package types

type Monitoring_SloWindowsBasedSliGoodTotalRatioThresholdPerformanceDistributionCutRange struct {
	/*
	   max value for the range (inclusive). If not given,
	   will be set to "infinity", defining an open range
	   ">= range.min"
	*/
	Max float64 `json:"max,omitempty" yaml:"max,omitempty"`

	/*
	   Min value for the range (inclusive). If not given,
	   will be set to "-infinity", defining an open range
	   "< range.max"
	*/
	Min float64 `json:"min,omitempty" yaml:"min,omitempty"`
}
