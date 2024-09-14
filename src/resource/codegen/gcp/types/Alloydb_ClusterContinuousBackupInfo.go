package types

type Alloydb_ClusterContinuousBackupInfo struct {
	/*
	   (Output)
	   The earliest restorable time that can be restored to. Output only field.
	*/
	EarliestRestorableTime string `json:"earliestRestorableTime,omitempty" yaml:"earliestRestorableTime,omitempty"`

	/*
	   (Output)
	   When ContinuousBackup was most recently enabled. Set to null if ContinuousBackup is not enabled.
	*/
	EnabledTime string `json:"enabledTime,omitempty" yaml:"enabledTime,omitempty"`

	/*
	   (Output)
	   Output only. The encryption information for the WALs and backups required for ContinuousBackup.
	   Structure is documented below.
	*/
	EncryptionInfos []Alloydb_ClusterContinuousBackupInfoEncryptionInfo `json:"encryptionInfos,omitempty" yaml:"encryptionInfos,omitempty"`

	/*
	   (Output)
	   Days of the week on which a continuous backup is taken. Output only field. Ignored if passed into the request.
	*/
	Schedules []string `json:"schedules,omitempty" yaml:"schedules,omitempty"`
}
