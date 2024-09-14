package types

type Netapp_VolumeSnapshotPolicyMonthlySchedule struct {
	// Set the minute of the hour to create the snapshot (0-59), defaults to the top of the hour (0).
	Minute int `json:"minute,omitempty" yaml:"minute,omitempty"`

	// The maximum number of snapshots to keep for the monthly schedule
	SnapshotsToKeep int `json:"snapshotsToKeep,omitempty" yaml:"snapshotsToKeep,omitempty"`

	// Set the day or days of the month to make a snapshot (1-31). Accepts a comma separated number of days. Defaults to '1'.
	DaysOfMonth string `json:"daysOfMonth,omitempty" yaml:"daysOfMonth,omitempty"`

	// Set the hour to create the snapshot (0-23), defaults to midnight (0).
	Hour int `json:"hour,omitempty" yaml:"hour,omitempty"`
}
