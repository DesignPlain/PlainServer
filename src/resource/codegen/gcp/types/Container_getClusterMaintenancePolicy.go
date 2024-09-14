package types

type Container_getClusterMaintenancePolicy struct {
	// Time window specified for daily maintenance operations. Specify start_time in RFC3339 format "HH:MM”, where HH : [00-23] and MM : [00-59] GMT.
	DailyMaintenanceWindows []Container_getClusterMaintenancePolicyDailyMaintenanceWindow `json:"dailyMaintenanceWindows,omitempty" yaml:"dailyMaintenanceWindows,omitempty"`

	// Exceptions to maintenance window. Non-emergency maintenance should not occur in these windows.
	MaintenanceExclusions []Container_getClusterMaintenancePolicyMaintenanceExclusion `json:"maintenanceExclusions,omitempty" yaml:"maintenanceExclusions,omitempty"`

	// Time window for recurring maintenance operations.
	RecurringWindows []Container_getClusterMaintenancePolicyRecurringWindow `json:"recurringWindows,omitempty" yaml:"recurringWindows,omitempty"`
}
