package types

type Efs_FileSystemProtection struct {
	// Indicates whether replication overwrite protection is enabled. Valid values: `ENABLED` or `DISABLED`.
	ReplicationOverwrite string `json:"replicationOverwrite,omitempty" yaml:"replicationOverwrite,omitempty"`
}
