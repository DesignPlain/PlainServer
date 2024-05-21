package types

type Netapp_VolumeRestoreParameters struct {
	/*
	   Full name of the snapshot to use for creating this volume.
	   `source_snapshot` and `source_backup` cannot be used simultaneously.
	   Format: `projects/{{project}}/locations/{{location}}/backupVaults/{{backupVaultId}}/backups/{{backup}}`.
	*/
	SourceBackup string `json:"sourceBackup,omitempty" yaml:"sourceBackup,omitempty"`

	/*
	   Full name of the snapshot to use for creating this volume.
	   `source_snapshot` and `source_backup` cannot be used simultaneously.
	   Format: `projects/{{project}}/locations/{{location}}/volumes/{{volume}}/snapshots/{{snapshot}}`.
	*/
	SourceSnapshot string `json:"sourceSnapshot,omitempty" yaml:"sourceSnapshot,omitempty"`
}
