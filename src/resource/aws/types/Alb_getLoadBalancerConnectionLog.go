package types

type Alb_getLoadBalancerConnectionLog struct {
	//
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	//
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
