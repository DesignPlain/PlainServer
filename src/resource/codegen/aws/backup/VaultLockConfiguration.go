package backup

type VaultLockConfiguration struct {
	// The minimum retention period that the vault retains its recovery points.
	MinRetentionDays int `json:"minRetentionDays,omitempty" yaml:"minRetentionDays,omitempty"`

	// Name of the backup vault to add a lock configuration for.
	BackupVaultName string `json:"backupVaultName,omitempty" yaml:"backupVaultName,omitempty"`

	// The number of days before the lock date. If omitted creates a vault lock in `governance` mode, otherwise it will create a vault lock in `compliance` mode.
	ChangeableForDays int `json:"changeableForDays,omitempty" yaml:"changeableForDays,omitempty"`

	// The maximum retention period that the vault retains its recovery points.
	MaxRetentionDays int `json:"maxRetentionDays,omitempty" yaml:"maxRetentionDays,omitempty"`
}
