package types



type Dataloss_PreventionJobTriggerTrigger struct {
	// For use with hybrid jobs. Jobs must be manually created and finished.
	Manual Dataloss_PreventionJobTriggerTriggerManual `json:"manual,omitempty" yaml:"manual,omitempty"`

	/*
	   Schedule for triggered jobs
	   Structure is documented below.
	*/
	Schedule Dataloss_PreventionJobTriggerTriggerSchedule `json:"schedule,omitempty" yaml:"schedule,omitempty"`
}
