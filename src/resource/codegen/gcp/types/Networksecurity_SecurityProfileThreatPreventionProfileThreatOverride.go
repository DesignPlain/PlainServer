package types

type Networksecurity_SecurityProfileThreatPreventionProfileThreatOverride struct {
	/*
	   Threat action.
	   Possible values are: `ALERT`, `ALLOW`, `DEFAULT_ACTION`, `DENY`.
	*/
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// Vendor-specific ID of a threat to override.
	ThreatId string `json:"threatId,omitempty" yaml:"threatId,omitempty"`

	/*
	   (Output)
	   Type of threat.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
