package types

type Rds_InstanceRestoreToPointInTime struct {
	// The date and time to restore from. Value must be a time in Universal Coordinated Time (UTC) format and must be before the latest restorable time for the DB instance. Cannot be specified with `use_latest_restorable_time`.
	RestoreTime string `json:"restoreTime,omitempty" yaml:"restoreTime,omitempty"`

	// The ARN of the automated backup from which to restore. Required if `source_db_instance_identifier` or `source_dbi_resource_id` is not specified.
	SourceDbInstanceAutomatedBackupsArn string `json:"sourceDbInstanceAutomatedBackupsArn,omitempty" yaml:"sourceDbInstanceAutomatedBackupsArn,omitempty"`

	// The identifier of the source DB instance from which to restore. Must match the identifier of an existing DB instance. Required if `source_db_instance_automated_backups_arn` or `source_dbi_resource_id` is not specified.
	SourceDbInstanceIdentifier string `json:"sourceDbInstanceIdentifier,omitempty" yaml:"sourceDbInstanceIdentifier,omitempty"`

	// The resource ID of the source DB instance from which to restore. Required if `source_db_instance_identifier` or `source_db_instance_automated_backups_arn` is not specified.
	SourceDbiResourceId string `json:"sourceDbiResourceId,omitempty" yaml:"sourceDbiResourceId,omitempty"`

	// A boolean value that indicates whether the DB instance is restored from the latest backup time. Defaults to `false`. Cannot be specified with `restore_time`.
	UseLatestRestorableTime bool `json:"useLatestRestorableTime,omitempty" yaml:"useLatestRestorableTime,omitempty"`
}
