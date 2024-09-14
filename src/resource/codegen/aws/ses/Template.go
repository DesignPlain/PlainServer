package ses

type Template struct {
	// The HTML body of the email. Must be less than 500KB in size, including both the text and HTML parts.
	Html string `json:"html,omitempty" yaml:"html,omitempty"`

	// The name of the template. Cannot exceed 64 characters. You will refer to this name when you send email.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The subject line of the email.
	Subject string `json:"subject,omitempty" yaml:"subject,omitempty"`

	// The email body that will be visible to recipients whose email clients do not display HTML. Must be less than 500KB in size, including both the text and HTML parts.
	Text string `json:"text,omitempty" yaml:"text,omitempty"`
}
