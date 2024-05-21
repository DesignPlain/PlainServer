package types

type Compute_ResourcePolicySnapshotSchedulePolicy struct {
	/*
	   Retention policy applied to snapshots created by this resource policy.
	   Structure is documented below.
	*/
	RetentionPolicy Compute_ResourcePolicySnapshotSchedulePolicyRetentionPolicy `json:"retentionPolicy,omitempty" yaml:"retentionPolicy,omitempty"`

	/*
	   Contains one of an `hourlySchedule`, `dailySchedule`, or `weeklySchedule`.
	   Structure is documented below.
	*/
	Schedule Compute_ResourcePolicySnapshotSchedulePolicySchedule `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	/*
	   Properties with which the snapshots are created, such as labels.
	   Structure is documented below.
	*/
	SnapshotProperties Compute_ResourcePolicySnapshotSchedulePolicySnapshotProperties `json:"snapshotProperties,omitempty" yaml:"snapshotProperties,omitempty"`
}
