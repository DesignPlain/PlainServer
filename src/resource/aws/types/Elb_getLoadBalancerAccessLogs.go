package types

type Elb_getLoadBalancerAccessLogs struct {
	//
	BucketPrefix string `json:"bucketPrefix,omitempty" yaml:"bucketPrefix,omitempty"`

	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	//
	Interval int `json:"interval,omitempty" yaml:"interval,omitempty"`

	//
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`
}
