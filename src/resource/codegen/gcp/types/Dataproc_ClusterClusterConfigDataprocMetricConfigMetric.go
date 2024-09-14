package types

type Dataproc_ClusterClusterConfigDataprocMetricConfigMetric struct {
	/*
	   One or more [available OSS metrics] (https://cloud.google.com/dataproc/docs/guides/monitoring#available_oss_metrics) to collect for the metric course.

	   - - -
	*/
	MetricOverrides []string `json:"metricOverrides,omitempty" yaml:"metricOverrides,omitempty"`

	// A source for the collection of Dataproc OSS metrics (see [available OSS metrics](https://cloud.google.com//dataproc/docs/guides/monitoring#available_oss_metrics)).
	MetricSource string `json:"metricSource,omitempty" yaml:"metricSource,omitempty"`
}
