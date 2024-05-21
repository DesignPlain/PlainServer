package types

type Sql_DatabaseInstanceSettingsBackupConfiguration struct {
	// The region where the backup will be stored
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// True if Point-in-time recovery is enabled. Will restart database if enabled after instance creation. Valid only for PostgreSQL and SQL Server instances.
	PointInTimeRecoveryEnabled bool `json:"pointInTimeRecoveryEnabled,omitempty" yaml:"pointInTimeRecoveryEnabled,omitempty"`

	/*
	   `HH:MM` format time indicating when backup
	   configuration starts.
	*/
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`

	// The number of days of transaction logs we retain for point in time restore, from 1-7. For PostgreSQL Enterprise Plus instances, the number of days of retained transaction logs can be set from 1 to 35.
	TransactionLogRetentionDays int `json:"transactionLogRetentionDays,omitempty" yaml:"transactionLogRetentionDays,omitempty"`

	// Backup retention settings. The configuration is detailed below.
	BackupRetentionSettings Sql_DatabaseInstanceSettingsBackupConfigurationBackupRetentionSettings `json:"backupRetentionSettings,omitempty" yaml:"backupRetentionSettings,omitempty"`

	/*
	   True if binary logging is enabled.
	   Can only be used with MySQL.
	*/
	BinaryLogEnabled bool `json:"binaryLogEnabled,omitempty" yaml:"binaryLogEnabled,omitempty"`

	// True if backup configuration is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
