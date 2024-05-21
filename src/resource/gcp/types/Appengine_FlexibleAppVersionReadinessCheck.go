package types

type Appengine_FlexibleAppVersionReadinessCheck struct {
	// Host header to send when performing a HTTP Readiness check. Example: "myapp.appspot.com"
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	// The request path.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// Number of consecutive successful checks required before receiving traffic. Default: 2.
	SuccessThreshold float64 `json:"successThreshold,omitempty" yaml:"successThreshold,omitempty"`

	// Time before the check is considered failed. Default: "4s"
	Timeout string `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	/*
	   A maximum time limit on application initialization, measured from moment the application successfully
	   replies to a healthcheck until it is ready to serve traffic. Default: "300s"
	*/
	AppStartTimeout string `json:"appStartTimeout,omitempty" yaml:"appStartTimeout,omitempty"`

	// Interval between health checks.  Default: "5s".
	CheckInterval string `json:"checkInterval,omitempty" yaml:"checkInterval,omitempty"`

	// Number of consecutive failed checks required before removing traffic. Default: 2.
	FailureThreshold float64 `json:"failureThreshold,omitempty" yaml:"failureThreshold,omitempty"`
}
