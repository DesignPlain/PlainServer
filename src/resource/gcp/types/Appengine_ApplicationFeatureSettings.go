package types

type Appengine_ApplicationFeatureSettings struct {
	/*
	   Set to false to use the legacy health check instead of the readiness
	   and liveness checks.
	*/
	SplitHealthChecks bool `json:"splitHealthChecks,omitempty" yaml:"splitHealthChecks,omitempty"`
}
