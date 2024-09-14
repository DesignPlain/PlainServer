package types

type Bigquery_JobStatusErrorResult struct {
	// The geographic location of the job. The default value is US.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// A human-readable description of the error.
	Message string `json:"message,omitempty" yaml:"message,omitempty"`

	// A short error code that summarizes the error.
	Reason string `json:"reason,omitempty" yaml:"reason,omitempty"`
}
