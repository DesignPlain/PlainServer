package types

type Cloudwatch_MetricStreamStatisticsConfiguration struct {
	// An array that defines the metrics that are to have additional statistics streamed. See details below.
	IncludeMetrics []Cloudwatch_MetricStreamStatisticsConfigurationIncludeMetric `json:"includeMetrics,omitempty" yaml:"includeMetrics,omitempty"`

	// The additional statistics to stream for the metrics listed in `include_metrics`.
	AdditionalStatistics []string `json:"additionalStatistics,omitempty" yaml:"additionalStatistics,omitempty"`
}
