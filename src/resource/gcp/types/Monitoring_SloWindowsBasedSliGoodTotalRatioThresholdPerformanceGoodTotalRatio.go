package types

type Monitoring_SloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatio struct {
	/*
	   A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
	   quantifying good service provided. Exactly two of
	   good, bad, or total service filter must be defined (where
	   good + bad = total is assumed)
	   Must have ValueType = DOUBLE or ValueType = INT64 and
	   must have MetricKind = DELTA or MetricKind = CUMULATIVE.
	*/
	GoodServiceFilter string `json:"goodServiceFilter,omitempty" yaml:"goodServiceFilter,omitempty"`

	/*
	   A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
	   quantifying total demanded service. Exactly two of
	   good, bad, or total service filter must be defined (where
	   good + bad = total is assumed)
	   Must have ValueType = DOUBLE or ValueType = INT64 and
	   must have MetricKind = DELTA or MetricKind = CUMULATIVE.
	*/
	TotalServiceFilter string `json:"totalServiceFilter,omitempty" yaml:"totalServiceFilter,omitempty"`

	/*
	   A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
	   quantifying bad service provided, either demanded service that
	   was not provided or demanded service that was of inadequate
	   quality. Exactly two of
	   good, bad, or total service filter must be defined (where
	   good + bad = total is assumed)
	   Must have ValueType = DOUBLE or ValueType = INT64 and
	   must have MetricKind = DELTA or MetricKind = CUMULATIVE.
	*/
	BadServiceFilter string `json:"badServiceFilter,omitempty" yaml:"badServiceFilter,omitempty"`
}
