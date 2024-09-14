package sesv2

type EmailIdentityFeedbackAttributes struct {
	// Sets the feedback forwarding configuration for the identity.
	EmailForwardingEnabled bool `json:"emailForwardingEnabled,omitempty" yaml:"emailForwardingEnabled,omitempty"`

	// The email identity.
	EmailIdentity string `json:"emailIdentity,omitempty" yaml:"emailIdentity,omitempty"`
}
