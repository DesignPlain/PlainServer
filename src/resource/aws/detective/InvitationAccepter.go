package detective

type InvitationAccepter struct {
	// ARN of the behavior graph that the member account is accepting the invitation for.
	GraphArn string `json:"graphArn,omitempty" yaml:"graphArn,omitempty"`
}
