package ssoadmin

type ApplicationAssignmentConfiguration struct {
	// Indicates whether users must have an explicit assignment to access the application. If `false`, all users have access to the application.
	AssignmentRequired bool `json:"assignmentRequired,omitempty" yaml:"assignmentRequired,omitempty"`

	// ARN of the application.
	ApplicationArn string `json:"applicationArn,omitempty" yaml:"applicationArn,omitempty"`
}
