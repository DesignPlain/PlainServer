package types

type Sql_getDatabaseInstanceSettingBackupConfiguration struct {
	// Location of the backup configuration.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// True if Point-in-time recovery is enabled.
	PointInTimeRecoveryEnabled bool `json:"pointInTimeRecoveryEnabled,omitempty" yaml:"pointInTimeRecoveryEnabled,omitempty"`

	// HH:MM format time indicating when backup configuration starts.
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`

	// The number of days of transaction logs we retain for point in time restore, from 1-7. (For PostgreSQL Enterprise Plus instances, from 1 to 35.)
	TransactionLogRetentionDays int `json:"transactionLogRetentionDays,omitempty" yaml:"transactionLogRetentionDays,omitempty"`

	//
	BackupRetentionSettings []Sql_getDatabaseInstanceSettingBackupConfigurationBackupRetentionSetting `json:"backupRetentionSettings,omitempty" yaml:"backupRetentionSettings,omitempty"`

	// True if binary logging is enabled. If settings.backup_configuration.enabled is false, this must be as well. Can only be used with MySQL.
	BinaryLogEnabled bool `json:"binaryLogEnabled,omitempty" yaml:"binaryLogEnabled,omitempty"`

	// True if backup configuration is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
