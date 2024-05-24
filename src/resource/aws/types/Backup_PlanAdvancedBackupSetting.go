package types

type Backup_PlanAdvancedBackupSetting struct {
	// Specifies the backup option for a selected resource. This option is only available for Windows VSS backup jobs. Set to `{ WindowsVSS = "enabled" }` to enable Windows VSS backup option and create a VSS Windows backup.
	BackupOptions map[string]string `json:"backupOptions,omitempty" yaml:"backupOptions,omitempty"`

	// The type of AWS resource to be backed up. For VSS Windows backups, the only supported resource type is Amazon EC2. Valid values: `EC2`.
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`
}
