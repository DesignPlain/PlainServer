package bigquery

type ReservationAssignment struct {
	// The resource which will use the reservation. E.g. projects/myproject, folders/123, organizations/456.
	Assignee string `json:"assignee,omitempty" yaml:"assignee,omitempty"`

	// Types of job, which could be specified when using the reservation. Possible values: JOB_TYPE_UNSPECIFIED, PIPELINE, QUERY
	JobType string `json:"jobType,omitempty" yaml:"jobType,omitempty"`

	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The reservation for the resource



	   - - -
	*/
	Reservation string `json:"reservation,omitempty" yaml:"reservation,omitempty"`
}
