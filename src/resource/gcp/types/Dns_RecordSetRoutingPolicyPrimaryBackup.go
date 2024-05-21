package types

type Dns_RecordSetRoutingPolicyPrimaryBackup struct {
	// Specifies whether to enable fencing for backup geo queries.
	EnableGeoFencingForBackups bool `json:"enableGeoFencingForBackups,omitempty" yaml:"enableGeoFencingForBackups,omitempty"`

	/*
	   The list of global primary targets to be health checked.
	   Structure is document below.
	*/
	Primary Dns_RecordSetRoutingPolicyPrimaryBackupPrimary `json:"primary,omitempty" yaml:"primary,omitempty"`

	// Specifies the percentage of traffic to send to the backup targets even when the primary targets are healthy.
	TrickleRatio float64 `json:"trickleRatio,omitempty" yaml:"trickleRatio,omitempty"`

	/*
	   The backup geo targets, which provide a regional failover policy for the otherwise global primary targets.
	   Structure is document above.
	*/
	BackupGeos []Dns_RecordSetRoutingPolicyPrimaryBackupBackupGeo `json:"backupGeos,omitempty" yaml:"backupGeos,omitempty"`
}
