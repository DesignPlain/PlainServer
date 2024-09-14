package types

type Sql_DatabaseInstanceRestoreBackupContext struct {
	// The ID of the backup run to restore from.
	BackupRunId int `json:"backupRunId,omitempty" yaml:"backupRunId,omitempty"`

	/*
	   The ID of the instance that the backup was taken from. If left empty,
	   this instance's ID will be used.
	*/
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	// The full project ID of the source instance.`
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
