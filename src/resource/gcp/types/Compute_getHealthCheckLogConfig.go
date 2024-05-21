package types

type Compute_getHealthCheckLogConfig struct {
	/*
	   Indicates whether or not to export logs. This is false by default,
	   which means no health check logging will be done.
	*/
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`
}
