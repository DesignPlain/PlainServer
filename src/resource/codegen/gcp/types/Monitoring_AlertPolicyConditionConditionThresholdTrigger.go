package types

type Monitoring_AlertPolicyConditionConditionThresholdTrigger struct {
	/*
	   The absolute number of time series
	   that must fail the predicate for the
	   condition to be triggered.
	*/
	Count int `json:"count,omitempty" yaml:"count,omitempty"`

	/*
	   The percentage of time series that
	   must fail the predicate for the
	   condition to be triggered.
	*/
	Percent float64 `json:"percent,omitempty" yaml:"percent,omitempty"`
}
