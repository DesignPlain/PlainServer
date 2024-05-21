package types

type Netapp_VolumeSnapshotPolicyWeeklySchedule struct {
	// Set the day or days of the week to make a snapshot. Accepts a comma separated days of the week. Defaults to 'Sunday'.
	Day string `json:"day,omitempty" yaml:"day,omitempty"`

	// Set the hour to create the snapshot (0-23), defaults to midnight (0).
	Hour int `json:"hour,omitempty" yaml:"hour,omitempty"`

	// Set the minute of the hour to create the snapshot (0-59), defaults to the top of the hour (0).
	Minute int `json:"minute,omitempty" yaml:"minute,omitempty"`

	// The maximum number of snapshots to keep for the weekly schedule.
	SnapshotsToKeep int `json:"snapshotsToKeep,omitempty" yaml:"snapshotsToKeep,omitempty"`
}
