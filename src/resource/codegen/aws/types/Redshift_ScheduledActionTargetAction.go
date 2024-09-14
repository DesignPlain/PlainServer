package types

type Redshift_ScheduledActionTargetAction struct {
	// An action that runs a `PauseCluster` API operation. Documented below.
	PauseCluster Redshift_ScheduledActionTargetActionPauseCluster `json:"pauseCluster,omitempty" yaml:"pauseCluster,omitempty"`

	// An action that runs a `ResizeCluster` API operation. Documented below.
	ResizeCluster Redshift_ScheduledActionTargetActionResizeCluster `json:"resizeCluster,omitempty" yaml:"resizeCluster,omitempty"`

	// An action that runs a `ResumeCluster` API operation. Documented below.
	ResumeCluster Redshift_ScheduledActionTargetActionResumeCluster `json:"resumeCluster,omitempty" yaml:"resumeCluster,omitempty"`
}
