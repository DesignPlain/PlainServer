package sesv2

type EmailIdentityMailFromAttributes struct {
	// The action to take if the required MX record isn't found when you send an email. Valid values: `USE_DEFAULT_VALUE`, `REJECT_MESSAGE`.
	BehaviorOnMxFailure string `json:"behaviorOnMxFailure,omitempty" yaml:"behaviorOnMxFailure,omitempty"`

	// The verified email identity.
	EmailIdentity string `json:"emailIdentity,omitempty" yaml:"emailIdentity,omitempty"`

	// The custom MAIL FROM domain that you want the verified identity to use. Required if `behavior_on_mx_failure` is `REJECT_MESSAGE`.
	MailFromDomain string `json:"mailFromDomain,omitempty" yaml:"mailFromDomain,omitempty"`
}
