package types

type Bigquery_JobStatus struct {
	/*
	   (Output)
	   Running state of the job. Valid states include 'PENDING', 'RUNNING', and 'DONE'.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	/*
	   (Output)
	   Final error result of the job. If present, indicates that the job has completed and was unsuccessful.
	   Structure is documented below.
	*/
	ErrorResults []Bigquery_JobStatusErrorResult `json:"errorResults,omitempty" yaml:"errorResults,omitempty"`

	/*
	   (Output)
	   The first errors encountered during the running of the job. The final message
	   includes the number of errors that caused the process to stop. Errors here do
	   not necessarily mean that the job has not completed or was unsuccessful.
	   Structure is documented below.
	*/
	Errors []Bigquery_JobStatusError `json:"errors,omitempty" yaml:"errors,omitempty"`
}
