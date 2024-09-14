package types

type Lb_LoadBalancerAccessLogs struct {
	// Boolean to enable / disable `access_logs`. Defaults to `false`, even when `bucket` is specified.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// S3 bucket prefix. Logs are stored in the root if not configured.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// S3 bucket name to store the logs in.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`
}
