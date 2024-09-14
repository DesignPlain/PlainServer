package types

type Sql_getDatabaseInstancesInstanceSettingBackupConfigurationBackupRetentionSetting struct {
	// Number of backups to retain.
	RetainedBackups int `json:"retainedBackups,omitempty" yaml:"retainedBackups,omitempty"`

	// The unit that 'retainedBackups' represents. Defaults to COUNT
	RetentionUnit string `json:"retentionUnit,omitempty" yaml:"retentionUnit,omitempty"`
}
