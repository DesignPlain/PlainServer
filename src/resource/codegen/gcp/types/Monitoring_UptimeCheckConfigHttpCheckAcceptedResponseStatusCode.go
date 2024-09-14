package types

type Monitoring_UptimeCheckConfigHttpCheckAcceptedResponseStatusCode struct {
	/*
	   A class of status codes to accept.
	   Possible values are: `STATUS_CLASS_1XX`, `STATUS_CLASS_2XX`, `STATUS_CLASS_3XX`, `STATUS_CLASS_4XX`, `STATUS_CLASS_5XX`, `STATUS_CLASS_ANY`.
	*/
	StatusClass string `json:"statusClass,omitempty" yaml:"statusClass,omitempty"`

	// A status code to accept.
	StatusValue int `json:"statusValue,omitempty" yaml:"statusValue,omitempty"`
}
