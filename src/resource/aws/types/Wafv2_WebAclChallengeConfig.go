package types

type Wafv2_WebAclChallengeConfig struct {
	// Defines custom immunity time. See `immunity_time_property` below for details.
	ImmunityTimeProperty Wafv2_WebAclChallengeConfigImmunityTimeProperty `json:"immunityTimeProperty,omitempty" yaml:"immunityTimeProperty,omitempty"`
}
