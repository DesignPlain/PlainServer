package types

type Cloudrunv2_getJobLatestCreatedExecution struct {
	/*
	   Completion timestamp of the execution.

	   A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	*/
	CompletionTime string `json:"completionTime,omitempty" yaml:"completionTime,omitempty"`

	/*
	   Creation timestamp of the execution.

	   A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	*/
	CreateTime string `json:"createTime,omitempty" yaml:"createTime,omitempty"`

	// The name of the Cloud Run v2 Job.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
