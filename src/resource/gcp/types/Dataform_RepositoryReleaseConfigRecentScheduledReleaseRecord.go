package types

type Dataform_RepositoryReleaseConfigRecentScheduledReleaseRecord struct {
	/*
	   (Output)
	   The error status encountered upon this attempt to create the compilation result, if the attempt was unsuccessful.
	   Structure is documented below.
	*/
	ErrorStatuses []Dataform_RepositoryReleaseConfigRecentScheduledReleaseRecordErrorStatus `json:"errorStatuses,omitempty" yaml:"errorStatuses,omitempty"`

	/*
	   (Output)
	   The timestamp of this release attempt.
	*/
	ReleaseTime string `json:"releaseTime,omitempty" yaml:"releaseTime,omitempty"`

	/*
	   (Output)
	   The name of the created compilation result, if one was successfully created. Must be in the format projects/-/locations/-/repositories/-/compilationResults/-.
	*/
	CompilationResult string `json:"compilationResult,omitempty" yaml:"compilationResult,omitempty"`
}
