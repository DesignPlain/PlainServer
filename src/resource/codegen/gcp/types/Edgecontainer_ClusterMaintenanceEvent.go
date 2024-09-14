package types

type Edgecontainer_ClusterMaintenanceEvent struct {
	// The time that the window first starts.
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`

	/*
	   (Output)
	   Indicates the maintenance event state.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	/*
	   (Output)
	   Indicates the maintenance event type.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   (Output)
	   The time when the maintenance event message was updated.
	*/
	UpdateTime string `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`

	/*
	   (Output)
	   UUID of the maintenance event.
	*/
	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`

	/*
	   The time that the window ends. The end time must take place after the
	   start time.
	*/
	EndTime string `json:"endTime,omitempty" yaml:"endTime,omitempty"`

	/*
	   (Output)
	   The operation for running the maintenance event. Specified in the format
	   projects/-/locations/-/operations/-. If the maintenance event is split
	   into multiple operations (e.g. due to maintenance windows), the latest
	   one is recorded.
	*/
	Operation string `json:"operation,omitempty" yaml:"operation,omitempty"`

	/*
	   (Output)
	   The schedule of the maintenance event.
	*/
	Schedule string `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	// The target cluster version. For example: "1.5.0".
	TargetVersion string `json:"targetVersion,omitempty" yaml:"targetVersion,omitempty"`

	/*
	   (Output)
	   The time when the maintenance event request was created.
	*/
	CreateTime string `json:"createTime,omitempty" yaml:"createTime,omitempty"`
}
