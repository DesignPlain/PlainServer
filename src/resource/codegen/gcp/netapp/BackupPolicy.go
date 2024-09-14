package netapp

type BackupPolicy struct {
	// Name of the region for the policy to apply to.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Labels as key value pairs. Example: `{ "owner": "Bob", "department": "finance", "purpose": "testing" }`.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Number of monthly backups to keep. Note that the sum of daily, weekly and monthly backups should be greater than 1.
	MonthlyBackupLimit int `json:"monthlyBackupLimit,omitempty" yaml:"monthlyBackupLimit,omitempty"`

	/*
	   The name of the backup policy. Needs to be unique per location.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Number of weekly backups to keep. Note that the sum of daily, weekly and monthly backups should be greater than 1.
	WeeklyBackupLimit int `json:"weeklyBackupLimit,omitempty" yaml:"weeklyBackupLimit,omitempty"`

	// Number of daily backups to keep. Note that the minimum daily backup limit is 2.
	DailyBackupLimit int `json:"dailyBackupLimit,omitempty" yaml:"dailyBackupLimit,omitempty"`

	/*
	   If enabled, make backups automatically according to the schedules.
	   This will be applied to all volumes that have this policy attached and enforced on volume level.
	*/
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
