package redshift

type SnapshotCopy struct {
	// Number of days to retain newly copied snapshots in the destination AWS Region after they are copied from the source AWS Region. If the value is `-1`, the manual snapshot is retained indefinitely.
	ManualSnapshotRetentionPeriod int `json:"manualSnapshotRetentionPeriod,omitempty" yaml:"manualSnapshotRetentionPeriod,omitempty"`

	// Number of days to retain automated snapshots in the destination region after they are copied from the source region.
	RetentionPeriod int `json:"retentionPeriod,omitempty" yaml:"retentionPeriod,omitempty"`

	// Name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted cluster are copied to the destination region.
	SnapshotCopyGrantName string `json:"snapshotCopyGrantName,omitempty" yaml:"snapshotCopyGrantName,omitempty"`

	// Identifier of the source cluster.
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`

	/*
	   AWS Region to copy snapshots to.

	   The following arguments are optional:
	*/
	DestinationRegion string `json:"destinationRegion,omitempty" yaml:"destinationRegion,omitempty"`
}
