package types

type Networksecurity_SecurityProfileThreatPreventionProfileSeverityOverride struct {
	/*
	   Threat action override.
	   Possible values are: `ALERT`, `ALLOW`, `DEFAULT_ACTION`, `DENY`.
	*/
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	/*
	   Severity level to match.
	   Possible values are: `CRITICAL`, `HIGH`, `INFORMATIONAL`, `LOW`, `MEDIUM`.
	*/
	Severity string `json:"severity,omitempty" yaml:"severity,omitempty"`
}
