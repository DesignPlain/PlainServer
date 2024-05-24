package types

type Wafv2_WebAclChallengeConfigImmunityTimeProperty struct {
	// The amount of time, in seconds, that a CAPTCHA or challenge timestamp is considered valid by AWS WAF. The default setting is 300.
	ImmunityTime int `json:"immunityTime,omitempty" yaml:"immunityTime,omitempty"`
}
