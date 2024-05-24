package types

type Appmesh_VirtualNodeSpecListenerOutlierDetection struct {
	// Base amount of time for which a host is ejected.
	BaseEjectionDuration Appmesh_VirtualNodeSpecListenerOutlierDetectionBaseEjectionDuration `json:"baseEjectionDuration,omitempty" yaml:"baseEjectionDuration,omitempty"`

	// Time interval between ejection sweep analysis.
	Interval Appmesh_VirtualNodeSpecListenerOutlierDetectionInterval `json:"interval,omitempty" yaml:"interval,omitempty"`

	/*
	   Maximum percentage of hosts in load balancing pool for upstream service that can be ejected. Will eject at least one host regardless of the value.
	   Minimum value of `0`. Maximum value of `100`.
	*/
	MaxEjectionPercent int `json:"maxEjectionPercent,omitempty" yaml:"maxEjectionPercent,omitempty"`

	// Number of consecutive `5xx` errors required for ejection. Minimum value of `1`.
	MaxServerErrors int `json:"maxServerErrors,omitempty" yaml:"maxServerErrors,omitempty"`
}
