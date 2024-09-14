package types

type Appstream_FleetComputeCapacity struct {
	// Total number of simultaneous streaming instances that are running.
	Running int `json:"running,omitempty" yaml:"running,omitempty"`

	// Number of currently available instances that can be used to stream sessions.
	Available int `json:"available,omitempty" yaml:"available,omitempty"`

	// Desired number of streaming instances.
	DesiredInstances int `json:"desiredInstances,omitempty" yaml:"desiredInstances,omitempty"`

	// Desired number of user sessions for a multi-session fleet. This is not allowed for single-session fleets.
	DesiredSessions int `json:"desiredSessions,omitempty" yaml:"desiredSessions,omitempty"`

	// Number of instances in use for streaming.
	InUse int `json:"inUse,omitempty" yaml:"inUse,omitempty"`
}
