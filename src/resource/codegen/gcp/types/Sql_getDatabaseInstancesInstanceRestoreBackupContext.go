package types

type Sql_getDatabaseInstancesInstanceRestoreBackupContext struct {
	// The ID of the instance that the backup was taken from.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	// The ID of the project in which the resources belong. If it is not provided, the provider project is used.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The ID of the backup run to restore from.
	BackupRunId int `json:"backupRunId,omitempty" yaml:"backupRunId,omitempty"`
}
