package types

type Edgecontainer_ClusterMaintenancePolicyWindowRecurringWindowWindow struct {
	/*
	   The time that the window ends. The end time must take place after the
	   start time.
	*/
	EndTime string `json:"endTime,omitempty" yaml:"endTime,omitempty"`

	// The time that the window first starts.
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`
}
