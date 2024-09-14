package types

type Gkebackup_BackupPlanBackupSchedule struct {
	/*
	   A standard cron string that defines a repeating schedule for
	   creating Backups via this BackupPlan.
	   If this is defined, then backupRetainDays must also be defined.
	*/
	CronSchedule string `json:"cronSchedule,omitempty" yaml:"cronSchedule,omitempty"`

	// This flag denotes whether automatic Backup creation is paused for this BackupPlan.
	Paused bool `json:"paused,omitempty" yaml:"paused,omitempty"`
}
