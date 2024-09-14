package types

type Opensearch_DomainAutoTuneOptions struct {
	// Whether to schedule Auto-Tune optimizations that require blue/green deployments during the domain's configured daily off-peak window. Defaults to `false`.
	UseOffPeakWindow bool `json:"useOffPeakWindow,omitempty" yaml:"useOffPeakWindow,omitempty"`

	// Auto-Tune desired state for the domain. Valid values: `ENABLED` or `DISABLED`.
	DesiredState string `json:"desiredState,omitempty" yaml:"desiredState,omitempty"`

	/*
	   Configuration block for Auto-Tune maintenance windows. Can be specified multiple times for each maintenance window. Detailed below.

	   --NOTE:-- Maintenance windows are deprecated and have been replaced with [off-peak windows](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/off-peak.html). Consequently, `maintenance_schedule` configuration blocks cannot be specified when `use_off_peak_window` is set to `true`.
	*/
	MaintenanceSchedules []Opensearch_DomainAutoTuneOptionsMaintenanceSchedule `json:"maintenanceSchedules,omitempty" yaml:"maintenanceSchedules,omitempty"`

	// Whether to roll back to default Auto-Tune settings when disabling Auto-Tune. Valid values: `DEFAULT_ROLLBACK` or `NO_ROLLBACK`.
	RollbackOnDisable string `json:"rollbackOnDisable,omitempty" yaml:"rollbackOnDisable,omitempty"`
}
