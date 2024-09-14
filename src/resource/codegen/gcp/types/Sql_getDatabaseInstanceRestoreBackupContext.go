package types

type Sql_getDatabaseInstanceRestoreBackupContext struct {
	// The ID of the backup run to restore from.
	BackupRunId int `json:"backupRunId,omitempty" yaml:"backupRunId,omitempty"`

	// The ID of the instance that the backup was taken from.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	// The ID of the project in which the resource belongs.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
