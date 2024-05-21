package types

type Appengine_FlexibleAppVersionLivenessCheck struct {
	// The request path.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// Number of consecutive successful checks required before considering the VM healthy. Default: 2.
	SuccessThreshold float64 `json:"successThreshold,omitempty" yaml:"successThreshold,omitempty"`

	// Time before the check is considered failed. Default: "4s"
	Timeout string `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	// Interval between health checks.
	CheckInterval string `json:"checkInterval,omitempty" yaml:"checkInterval,omitempty"`

	// Number of consecutive failed checks required before considering the VM unhealthy. Default: 4.
	FailureThreshold float64 `json:"failureThreshold,omitempty" yaml:"failureThreshold,omitempty"`

	// Host header to send when performing a HTTP Readiness check. Example: "myapp.appspot.com"
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	/*
	   The initial delay before starting to execute the checks. Default: "300s"

	   - - -
	*/
	InitialDelay string `json:"initialDelay,omitempty" yaml:"initialDelay,omitempty"`
}
