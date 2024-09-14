package firestore

import types "libds/gcp/types"

type BackupSchedule struct {
	// The Firestore database id. Defaults to `"(default)"`.
	Database string `json:"database,omitempty" yaml:"database,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days.
	   A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	   For a daily backup recurrence, set this to a value up to 7 days. If you set a weekly backup recurrence, set this to a value up to 14 weeks.


	   - - -
	*/
	Retention string `json:"retention,omitempty" yaml:"retention,omitempty"`

	/*
	   For a schedule that runs weekly on a specific day and time.
	   Structure is documented below.
	*/
	WeeklyRecurrence types.Firestore_BackupScheduleWeeklyRecurrence `json:"weeklyRecurrence,omitempty" yaml:"weeklyRecurrence,omitempty"`

	// For a schedule that runs daily at a specified time.
	DailyRecurrence types.Firestore_BackupScheduleDailyRecurrence `json:"dailyRecurrence,omitempty" yaml:"dailyRecurrence,omitempty"`
}
