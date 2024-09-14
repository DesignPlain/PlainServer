package types

type Dataplex_TaskExecutionStatus struct {
	/*
	   (Output)
	   latest job execution.
	   Structure is documented below.
	*/
	LatestJobs []Dataplex_TaskExecutionStatusLatestJob `json:"latestJobs,omitempty" yaml:"latestJobs,omitempty"`

	/*
	   (Output)
	   Last update time of the status.
	*/
	UpdateTime string `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`
}
