package types

type Glue_JobExecutionProperty struct {
	// The maximum number of concurrent runs allowed for a job. The default is 1.
	MaxConcurrentRuns int `json:"maxConcurrentRuns,omitempty" yaml:"maxConcurrentRuns,omitempty"`
}
