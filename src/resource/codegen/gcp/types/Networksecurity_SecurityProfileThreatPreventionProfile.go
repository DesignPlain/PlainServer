package types

type Networksecurity_SecurityProfileThreatPreventionProfile struct {
	/*
	   The configuration for overriding threats actions by severity match.
	   Structure is documented below.
	*/
	SeverityOverrides []Networksecurity_SecurityProfileThreatPreventionProfileSeverityOverride `json:"severityOverrides,omitempty" yaml:"severityOverrides,omitempty"`

	/*
	   The configuration for overriding threats actions by threat id match.
	   If a threat is matched both by configuration provided in severity overrides
	   and threat overrides, the threat overrides action is applied.
	   Structure is documented below.
	*/
	ThreatOverrides []Networksecurity_SecurityProfileThreatPreventionProfileThreatOverride `json:"threatOverrides,omitempty" yaml:"threatOverrides,omitempty"`
}
