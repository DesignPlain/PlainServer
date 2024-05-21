package types

type Sql_DatabaseInstanceSettingsSqlServerAuditConfig struct {
	// The name of the destination bucket (e.g., gs://mybucket).
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// How long to keep generated audit files. A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	RetentionInterval string `json:"retentionInterval,omitempty" yaml:"retentionInterval,omitempty"`

	// How often to upload generated audit files. A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	UploadInterval string `json:"uploadInterval,omitempty" yaml:"uploadInterval,omitempty"`
}
