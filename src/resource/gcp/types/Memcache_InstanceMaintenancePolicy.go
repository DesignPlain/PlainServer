package types

type Memcache_InstanceMaintenancePolicy struct {
	/*
	   Required. Maintenance window that is applied to resources covered by this policy.
	   Minimum 1. For the current version, the maximum number of weekly_maintenance_windows
	   is expected to be one.
	   Structure is documented below.
	*/
	WeeklyMaintenanceWindows []Memcache_InstanceMaintenancePolicyWeeklyMaintenanceWindow `json:"weeklyMaintenanceWindows,omitempty" yaml:"weeklyMaintenanceWindows,omitempty"`

	/*
	   (Output)
	   Output only. The time when the policy was created.
	   A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	   resolution and up to nine fractional digits
	*/
	CreateTime string `json:"createTime,omitempty" yaml:"createTime,omitempty"`

	/*
	   Optional. Description of what this policy is for.
	   Create/Update methods return INVALID_ARGUMENT if the
	   length is greater than 512.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   (Output)
	   Output only. The time when the policy was updated.
	   A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	   resolution and up to nine fractional digits.
	*/
	UpdateTime string `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`
}
