package types

type Alloydb_ClusterRestoreBackupSource struct {
	// The name of the backup that this cluster is restored from.
	BackupName string `json:"backupName,omitempty" yaml:"backupName,omitempty"`
}
