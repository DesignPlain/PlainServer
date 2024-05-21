package types

type Monitoring_SloWindowsBasedSliMetricSumInRange struct {
	/*
	   Range of numerical values. The computed good_service
	   will be the count of values x in the Distribution such
	   that range.min <= x <= range.max. inclusive of min and
	   max. Open ranges can be defined by setting
	   just one of min or max. Summed value `X` should satisfy
	   `range.min <= X <= range.max` for a good window.
	   Structure is documented below.
	*/
	Range Monitoring_SloWindowsBasedSliMetricSumInRangeRange `json:"range,omitempty" yaml:"range,omitempty"`

	/*
	   A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
	   specifying the TimeSeries to use for evaluating window
	   quality. The provided TimeSeries must have
	   ValueType = INT64 or ValueType = DOUBLE and
	   MetricKind = GAUGE.
	   Summed value `X` should satisfy
	   `range.min <= X <= range.max` for a good window.
	*/
	TimeSeries string `json:"timeSeries,omitempty" yaml:"timeSeries,omitempty"`
}
