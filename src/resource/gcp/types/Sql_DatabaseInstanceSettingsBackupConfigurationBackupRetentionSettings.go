package types

type Sql_DatabaseInstanceSettingsBackupConfigurationBackupRetentionSettings struct {
	// The unit that 'retained_backups' represents. Defaults to `COUNT`.
	RetentionUnit string `json:"retentionUnit,omitempty" yaml:"retentionUnit,omitempty"`

	/*
	   Depending on the value of retention_unit, this is used to determine if a backup needs to be deleted. If retention_unit
	   is 'COUNT', we will retain this many backups.
	*/
	RetainedBackups int `json:"retainedBackups,omitempty" yaml:"retainedBackups,omitempty"`
}
