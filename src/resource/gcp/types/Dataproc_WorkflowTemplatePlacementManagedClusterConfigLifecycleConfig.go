package types

type Dataproc_WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig struct {
	// The time when cluster will be auto-deleted (see JSON representation of (https://developers.google.com/protocol-buffers/docs/proto3#json)).
	AutoDeleteTime string `json:"autoDeleteTime,omitempty" yaml:"autoDeleteTime,omitempty"`

	// The lifetime duration of cluster. The cluster will be auto-deleted at the end of this period. Minimum value is 10 minutes; maximum value is 14 days (see JSON representation of (https://developers.google.com/protocol-buffers/docs/proto3#json)).
	AutoDeleteTtl string `json:"autoDeleteTtl,omitempty" yaml:"autoDeleteTtl,omitempty"`

	// The duration to keep the cluster alive while idling (when no jobs are running). Passing this threshold will cause the cluster to be deleted. Minimum value is 5 minutes; maximum value is 14 days (see JSON representation of (https://developers.google.com/protocol-buffers/docs/proto3#json).
	IdleDeleteTtl string `json:"idleDeleteTtl,omitempty" yaml:"idleDeleteTtl,omitempty"`

	// Output only. The time when cluster became idle (most recent job finished) and became eligible for deletion due to idleness (see JSON representation of (https://developers.google.com/protocol-buffers/docs/proto3#json)).
	IdleStartTime string `json:"idleStartTime,omitempty" yaml:"idleStartTime,omitempty"`
}
