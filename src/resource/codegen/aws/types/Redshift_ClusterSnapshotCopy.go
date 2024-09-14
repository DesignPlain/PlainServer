package types

type Redshift_ClusterSnapshotCopy struct {
	// The name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted cluster are copied to the destination region.
	GrantName string `json:"grantName,omitempty" yaml:"grantName,omitempty"`

	// The number of days to retain automated snapshots in the destination region after they are copied from the source region. Defaults to `7`.
	RetentionPeriod int `json:"retentionPeriod,omitempty" yaml:"retentionPeriod,omitempty"`

	// The destination region that you want to copy snapshots to.
	DestinationRegion string `json:"destinationRegion,omitempty" yaml:"destinationRegion,omitempty"`
}
