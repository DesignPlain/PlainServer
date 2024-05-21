package logging

import types "DesignSphere_Server/src/resource/gcp/types"

type FolderBucketConfig struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId string `json:"bucketId,omitempty" yaml:"bucketId,omitempty"`

	/*
	   The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
	   key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
	   updating the log bucket. Changing the KMS key is allowed.
	*/
	CmekSettings types.Logging_FolderBucketConfigCmekSettings `json:"cmekSettings,omitempty" yaml:"cmekSettings,omitempty"`

	// Describes this bucket.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The parent resource that contains the logging bucket.
	Folder string `json:"folder,omitempty" yaml:"folder,omitempty"`

	// A list of indexed fields and related configuration data. Structure is documented below.
	IndexConfigs []types.Logging_FolderBucketConfigIndexConfig `json:"indexConfigs,omitempty" yaml:"indexConfigs,omitempty"`

	// The location of the bucket.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
	RetentionDays int `json:"retentionDays,omitempty" yaml:"retentionDays,omitempty"`
}
