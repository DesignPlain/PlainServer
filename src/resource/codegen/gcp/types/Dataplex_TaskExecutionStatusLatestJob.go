package types

type Dataplex_TaskExecutionStatusLatestJob struct {
	/*
	   (Output)
	   The number of times the job has been retried (excluding the initial attempt).
	*/
	RetryCount int `json:"retryCount,omitempty" yaml:"retryCount,omitempty"`

	/*
	   (Output)
	   The full resource name for the job run under a particular service.
	*/
	ServiceJob string `json:"serviceJob,omitempty" yaml:"serviceJob,omitempty"`

	/*
	   (Output)
	   Execution state for the job.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	/*
	   (Output)
	   System generated globally unique ID for the job.
	*/
	Uid string `json:"uid,omitempty" yaml:"uid,omitempty"`

	/*
	   (Output)
	   The time when the job ended.
	*/
	EndTime string `json:"endTime,omitempty" yaml:"endTime,omitempty"`

	/*
	   (Output)
	   Additional information about the current state.
	*/
	Message string `json:"message,omitempty" yaml:"message,omitempty"`

	// The first run of the task will be after this time. If not specified, the task will run shortly after being submitted if ON_DEMAND and based on the schedule if RECURRING.
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`

	/*
	   (Output)
	   The relative resource name of the job, of the form: projects/{project_number}/locations/{locationId}/lakes/{lakeId}/tasks/{taskId}/jobs/{jobId}.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   (Output)
	   The underlying service running a job.
	*/
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}
