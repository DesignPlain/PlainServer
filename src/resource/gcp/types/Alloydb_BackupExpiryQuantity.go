package types

type Alloydb_BackupExpiryQuantity struct {
	/*
	   (Output)
	   Output only. The backup's position among its backups with the same source cluster and type, by descending chronological order create time (i.e. newest first).
	*/
	RetentionCount int `json:"retentionCount,omitempty" yaml:"retentionCount,omitempty"`

	/*
	   (Output)
	   Output only. The length of the quantity-based queue, specified by the backup's retention policy.
	*/
	TotalRetentionCount int `json:"totalRetentionCount,omitempty" yaml:"totalRetentionCount,omitempty"`
}
