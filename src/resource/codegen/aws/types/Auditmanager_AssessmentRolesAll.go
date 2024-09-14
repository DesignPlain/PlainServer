package types

type Auditmanager_AssessmentRolesAll struct {
	// Amazon Resource Name (ARN) of the IAM role.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Type of customer persona. For assessment creation, type must always be `PROCESS_OWNER`.
	RoleType string `json:"roleType,omitempty" yaml:"roleType,omitempty"`
}
