package types

type Netapp_VolumeSnapshotPolicy struct {
	/*
	   Daily schedule policy.
	   Structure is documented below.
	*/
	DailySchedule Netapp_VolumeSnapshotPolicyDailySchedule `json:"dailySchedule,omitempty" yaml:"dailySchedule,omitempty"`

	/*
	   Enables automated snapshot creation according to defined schedule. Default is false.
	   To disable automatic snapshot creation you have to remove the whole snapshot_policy block.
	*/
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	/*
	   Hourly schedule policy.
	   Structure is documented below.
	*/
	HourlySchedule Netapp_VolumeSnapshotPolicyHourlySchedule `json:"hourlySchedule,omitempty" yaml:"hourlySchedule,omitempty"`

	/*
	   Monthly schedule policy.
	   Structure is documented below.
	*/
	MonthlySchedule Netapp_VolumeSnapshotPolicyMonthlySchedule `json:"monthlySchedule,omitempty" yaml:"monthlySchedule,omitempty"`

	/*
	   Weekly schedule policy.
	   Structure is documented below.
	*/
	WeeklySchedule Netapp_VolumeSnapshotPolicyWeeklySchedule `json:"weeklySchedule,omitempty" yaml:"weeklySchedule,omitempty"`
}
