package types

type Redis_getInstanceMaintenancePolicy struct {
	/*
	   Output only. The time when the policy was created.
	   A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	   resolution and up to nine fractional digits.
	*/
	CreateTime string `json:"createTime,omitempty" yaml:"createTime,omitempty"`

	/*
	   Optional. Description of what this policy is for.
	   Create/Update methods return INVALID_ARGUMENT if the
	   length is greater than 512.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Output only. The time when the policy was last updated.
	   A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	   resolution and up to nine fractional digits.
	*/
	UpdateTime string `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`

	/*
	   Optional. Maintenance window that is applied to resources covered by this policy.
	   Minimum 1. For the current version, the maximum number
	   of weekly_window is expected to be one.
	*/
	WeeklyMaintenanceWindows []Redis_getInstanceMaintenancePolicyWeeklyMaintenanceWindow `json:"weeklyMaintenanceWindows,omitempty" yaml:"weeklyMaintenanceWindows,omitempty"`
}
