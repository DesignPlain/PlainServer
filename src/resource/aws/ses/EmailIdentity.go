package ses

type EmailIdentity struct {
	// The email address to assign to SES.
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}
