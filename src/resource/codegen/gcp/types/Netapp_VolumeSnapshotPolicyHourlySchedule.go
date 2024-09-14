package types

type Netapp_VolumeSnapshotPolicyHourlySchedule struct {
	// The maximum number of snapshots to keep for the hourly schedule.
	SnapshotsToKeep int `json:"snapshotsToKeep,omitempty" yaml:"snapshotsToKeep,omitempty"`

	// Set the minute of the hour to create the snapshot (0-59), defaults to the top of the hour (0).
	Minute int `json:"minute,omitempty" yaml:"minute,omitempty"`
}
