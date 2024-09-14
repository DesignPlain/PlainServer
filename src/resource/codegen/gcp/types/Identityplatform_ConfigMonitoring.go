package types

type Identityplatform_ConfigMonitoring struct {
	/*
	   Configuration for logging requests made to this project to Stackdriver Logging
	   Structure is documented below.
	*/
	RequestLogging Identityplatform_ConfigMonitoringRequestLogging `json:"requestLogging,omitempty" yaml:"requestLogging,omitempty"`
}
