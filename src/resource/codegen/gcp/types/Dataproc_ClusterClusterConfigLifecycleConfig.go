package types

type Dataproc_ClusterClusterConfigLifecycleConfig struct {
	/*
	   The time when cluster will be auto-deleted.
	   A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	   Example: "2014-10-02T15:01:23.045123456Z".

	   - - -
	*/
	AutoDeleteTime string `json:"autoDeleteTime,omitempty" yaml:"autoDeleteTime,omitempty"`

	/*
	   The duration to keep the cluster alive while idling
	   (no jobs running). After this TTL, the cluster will be deleted. Valid range: [10m, 14d].
	*/
	IdleDeleteTtl string `json:"idleDeleteTtl,omitempty" yaml:"idleDeleteTtl,omitempty"`

	// Time when the cluster became idle (most recent job finished) and became eligible for deletion due to idleness.
	IdleStartTime string `json:"idleStartTime,omitempty" yaml:"idleStartTime,omitempty"`
}
