package types

type Ssoadmin_getApplicationAssignmentsApplicationAssignment struct {
	// Entity type for which the assignment will be created. Valid values are `USER` or `GROUP`.
	PrincipalType string `json:"principalType,omitempty" yaml:"principalType,omitempty"`

	// ARN of the application.
	ApplicationArn string `json:"applicationArn,omitempty" yaml:"applicationArn,omitempty"`

	// An identifier for an object in IAM Identity Center, such as a user or group.
	PrincipalId string `json:"principalId,omitempty" yaml:"principalId,omitempty"`
}
