package types

type Monitoring_SloBasicSliLatency struct {
	/*
	   A duration string, e.g. 10s.
	   Good service is defined to be the count of requests made to
	   this service that return in no more than threshold.
	*/
	Threshold string `json:"threshold,omitempty" yaml:"threshold,omitempty"`
}
