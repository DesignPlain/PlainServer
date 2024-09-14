package logging

import types "libds/gcp/types"

type ProjectBucketConfig struct {
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays int `json:"retentionDays,omitempty" yaml:"retentionDays,omitempty"`

	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed. Structure is documented below.
	CmekSettings types.Logging_ProjectBucketConfigCmekSettings `json:"cmekSettings,omitempty" yaml:"cmekSettings,omitempty"`

	// Describes this bucket.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The location of the bucket.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Whether the bucket is locked. The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
	Locked bool `json:"locked,omitempty" yaml:"locked,omitempty"`

	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId string `json:"bucketId,omitempty" yaml:"bucketId,omitempty"`

	// Whether or not Log Analytics is enabled. Logs for buckets with Log Analytics enabled can be queried in the --Log Analytics-- page using SQL queries. Cannot be disabled once enabled.
	EnableAnalytics bool `json:"enableAnalytics,omitempty" yaml:"enableAnalytics,omitempty"`

	// A list of indexed fields and related configuration data. Structure is documented below.
	IndexConfigs []types.Logging_ProjectBucketConfigIndexConfig `json:"indexConfigs,omitempty" yaml:"indexConfigs,omitempty"`

	// The parent resource that contains the logging bucket.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
