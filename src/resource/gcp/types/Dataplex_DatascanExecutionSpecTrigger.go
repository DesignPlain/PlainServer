package types

type Dataplex_DatascanExecutionSpecTrigger struct {
	// The scan runs once via dataScans.run API.
	OnDemand Dataplex_DatascanExecutionSpecTriggerOnDemand `json:"onDemand,omitempty" yaml:"onDemand,omitempty"`

	/*
	   The scan is scheduled to run periodically.
	   Structure is documented below.
	*/
	Schedule Dataplex_DatascanExecutionSpecTriggerSchedule `json:"schedule,omitempty" yaml:"schedule,omitempty"`
}
