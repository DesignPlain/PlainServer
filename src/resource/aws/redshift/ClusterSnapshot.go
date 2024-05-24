package redshift

type ClusterSnapshot struct {
	// The cluster identifier for which you want a snapshot.
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`

	// The number of days that a manual snapshot is retained. If the value is `-1`, the manual snapshot is retained indefinitely. Valid values are -1 and between `1` and `3653`.
	ManualSnapshotRetentionPeriod int `json:"manualSnapshotRetentionPeriod,omitempty" yaml:"manualSnapshotRetentionPeriod,omitempty"`

	// A unique identifier for the snapshot that you are requesting. This identifier must be unique for all snapshots within the Amazon Web Services account.
	SnapshotIdentifier string `json:"snapshotIdentifier,omitempty" yaml:"snapshotIdentifier,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
