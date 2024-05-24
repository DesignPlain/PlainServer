package backup

type VaultNotifications struct {
	// An array of events that indicate the status of jobs to back up resources to the backup vault.
	BackupVaultEvents []string `json:"backupVaultEvents,omitempty" yaml:"backupVaultEvents,omitempty"`

	// Name of the backup vault to add notifications for.
	BackupVaultName string `json:"backupVaultName,omitempty" yaml:"backupVaultName,omitempty"`

	// The Amazon Resource Name (ARN) that specifies the topic for a backup vault’s events
	SnsTopicArn string `json:"snsTopicArn,omitempty" yaml:"snsTopicArn,omitempty"`
}
