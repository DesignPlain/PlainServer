package types

type Compute_getResourcePolicySnapshotSchedulePolicy struct {
	// Retention policy applied to snapshots created by this resource policy.
	RetentionPolicies []Compute_getResourcePolicySnapshotSchedulePolicyRetentionPolicy `json:"retentionPolicies,omitempty" yaml:"retentionPolicies,omitempty"`

	// Contains one of an 'hourlySchedule', 'dailySchedule', or 'weeklySchedule'.
	Schedules []Compute_getResourcePolicySnapshotSchedulePolicySchedule `json:"schedules,omitempty" yaml:"schedules,omitempty"`

	// Properties with which the snapshots are created, such as labels.
	SnapshotProperties []Compute_getResourcePolicySnapshotSchedulePolicySnapshotProperty `json:"snapshotProperties,omitempty" yaml:"snapshotProperties,omitempty"`
}
