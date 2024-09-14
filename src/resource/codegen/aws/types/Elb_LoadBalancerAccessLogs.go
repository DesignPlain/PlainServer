package types

type Elb_LoadBalancerAccessLogs struct {
	// The S3 bucket name to store the logs in.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// The S3 bucket prefix. Logs are stored in the root if not configured.
	BucketPrefix string `json:"bucketPrefix,omitempty" yaml:"bucketPrefix,omitempty"`

	// Boolean to enable / disable `access_logs`. Default is `true`
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The publishing interval in minutes. Valid values: `5` and `60`. Default: `60`
	Interval int `json:"interval,omitempty" yaml:"interval,omitempty"`
}
