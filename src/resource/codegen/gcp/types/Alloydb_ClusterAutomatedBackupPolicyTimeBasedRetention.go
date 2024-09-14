package types

type Alloydb_ClusterAutomatedBackupPolicyTimeBasedRetention struct {
	/*
	   The retention period.
	   A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	*/
	RetentionPeriod string `json:"retentionPeriod,omitempty" yaml:"retentionPeriod,omitempty"`
}
