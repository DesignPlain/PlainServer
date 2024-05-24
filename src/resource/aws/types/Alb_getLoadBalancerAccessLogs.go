package types

type Alb_getLoadBalancerAccessLogs struct {
	//
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	//
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
