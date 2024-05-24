package ses

type MailFrom struct {
	/*
	   Subdomain (of above domain) which is to be used as MAIL FROM address (Required for DMARC validation)

	   The following arguments are optional:
	*/
	MailFromDomain string `json:"mailFromDomain,omitempty" yaml:"mailFromDomain,omitempty"`

	// The action that you want Amazon SES to take if it cannot successfully read the required MX record when you send an email. Defaults to `UseDefaultValue`. See the [SES API documentation](https://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityMailFromDomain.html) for more information.
	BehaviorOnMxFailure string `json:"behaviorOnMxFailure,omitempty" yaml:"behaviorOnMxFailure,omitempty"`

	// Verified domain name or email identity to generate DKIM tokens for.
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`
}
