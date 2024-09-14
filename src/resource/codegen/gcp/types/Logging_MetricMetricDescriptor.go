package types

type Logging_MetricMetricDescriptor struct {
	/*
	   Whether the measurement is an integer, a floating-point number, etc.
	   Some combinations of metricKind and valueType might not be supported.
	   For counter metrics, set this to INT64.
	   Possible values are: `BOOL`, `INT64`, `DOUBLE`, `STRING`, `DISTRIBUTION`, `MONEY`.
	*/
	ValueType string `json:"valueType,omitempty" yaml:"valueType,omitempty"`

	/*
	   A concise name for the metric, which can be displayed in user interfaces. Use sentence case
	   without an ending period, for example "Request count". This field is optional but it is
	   recommended to be set for any metrics associated with user-visible concepts, such as Quota.
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The set of labels that can be used to describe a specific instance of this metric type. For
	   example, the appengine.googleapis.com/http/server/response_latencies metric type has a label
	   for the HTTP response code, response_code, so you can look at latencies for successful responses
	   or just for responses that failed.
	   Structure is documented below.
	*/
	Labels []Logging_MetricMetricDescriptorLabel `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Whether the metric records instantaneous values, changes to a value, etc.
	   Some combinations of metricKind and valueType might not be supported.
	   For counter metrics, set this to DELTA.
	   Possible values are: `DELTA`, `GAUGE`, `CUMULATIVE`.
	*/
	MetricKind string `json:"metricKind,omitempty" yaml:"metricKind,omitempty"`

	/*
	   The unit in which the metric value is reported. It is only applicable if the valueType is
	   `INT64`, `DOUBLE`, or `DISTRIBUTION`. The supported units are a subset of
	   [The Unified Code for Units of Measure](http://unitsofmeasure.org/ucum.html) standard
	*/
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`
}
