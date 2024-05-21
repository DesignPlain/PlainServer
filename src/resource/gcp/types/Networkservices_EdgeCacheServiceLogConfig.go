package types

type Networkservices_EdgeCacheServiceLogConfig struct {
	// Specifies whether to enable logging for traffic served by this service.
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`

	/*
	   Configures the sampling rate of requests, where 1.0 means all logged requests are reported and 0.0 means no logged requests are reported. The default value is 1.0, and the value of the field must be in [0, 1].
	   This field can only be specified if logging is enabled for this service.
	*/
	SampleRate float64 `json:"sampleRate,omitempty" yaml:"sampleRate,omitempty"`
}
