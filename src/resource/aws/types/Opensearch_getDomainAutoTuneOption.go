package types

type Opensearch_getDomainAutoTuneOption struct {
	// Auto-Tune desired state for the domain.
	DesiredState string `json:"desiredState,omitempty" yaml:"desiredState,omitempty"`

	// A list of the nested configurations for the Auto-Tune maintenance windows of the domain.
	MaintenanceSchedules []Opensearch_getDomainAutoTuneOptionMaintenanceSchedule `json:"maintenanceSchedules,omitempty" yaml:"maintenanceSchedules,omitempty"`

	// Whether the domain is set to roll back to default Auto-Tune settings when disabling Auto-Tune.
	RollbackOnDisable string `json:"rollbackOnDisable,omitempty" yaml:"rollbackOnDisable,omitempty"`
}
