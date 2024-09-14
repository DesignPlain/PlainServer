package types

type Elasticsearch_DomainAutoTuneOptions struct {
	// The Auto-Tune desired state for the domain. Valid values: `ENABLED` or `DISABLED`.
	DesiredState string `json:"desiredState,omitempty" yaml:"desiredState,omitempty"`

	// Configuration block for Auto-Tune maintenance windows. Can be specified multiple times for each maintenance window. Detailed below.
	MaintenanceSchedules []Elasticsearch_DomainAutoTuneOptionsMaintenanceSchedule `json:"maintenanceSchedules,omitempty" yaml:"maintenanceSchedules,omitempty"`

	// Whether to roll back to default Auto-Tune settings when disabling Auto-Tune. Valid values: `DEFAULT_ROLLBACK` or `NO_ROLLBACK`.
	RollbackOnDisable string `json:"rollbackOnDisable,omitempty" yaml:"rollbackOnDisable,omitempty"`
}
