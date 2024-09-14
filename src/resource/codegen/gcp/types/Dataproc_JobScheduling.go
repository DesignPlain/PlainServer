package types

type Dataproc_JobScheduling struct {
	// Maximum number of times in total a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed.
	MaxFailuresTotal int `json:"maxFailuresTotal,omitempty" yaml:"maxFailuresTotal,omitempty"`

	// Maximum number of times per hour a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed.
	MaxFailuresPerHour int `json:"maxFailuresPerHour,omitempty" yaml:"maxFailuresPerHour,omitempty"`
}
