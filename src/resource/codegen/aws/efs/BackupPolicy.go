package efs

import types "libds/aws/types"

type BackupPolicy struct {
	// A backup_policy object (documented below).
	BackupPolicy types.Efs_BackupPolicyBackupPolicy `json:"backupPolicy,omitempty" yaml:"backupPolicy,omitempty"`

	// The ID of the EFS file system.
	FileSystemId string `json:"fileSystemId,omitempty" yaml:"fileSystemId,omitempty"`
}
