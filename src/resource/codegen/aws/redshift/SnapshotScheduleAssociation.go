package redshift

type SnapshotScheduleAssociation struct {
	// The cluster identifier.
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`

	// The snapshot schedule identifier.
	ScheduleIdentifier string `json:"scheduleIdentifier,omitempty" yaml:"scheduleIdentifier,omitempty"`
}
