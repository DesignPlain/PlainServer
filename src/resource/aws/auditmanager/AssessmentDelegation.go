package auditmanager

type AssessmentDelegation struct {
	/*
	   Type of customer persona. For assessment delegation, type must always be `RESOURCE_OWNER`.

	   The following arguments are optional:
	*/
	RoleType string `json:"roleType,omitempty" yaml:"roleType,omitempty"`

	// Identifier for the assessment.
	AssessmentId string `json:"assessmentId,omitempty" yaml:"assessmentId,omitempty"`

	// Comment describing the delegation request.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// Assessment control set name. This value is the control set name used during assessment creation (not the AWS-generated ID). The `_id` suffix on this attribute has been preserved to be consistent with the underlying AWS API.
	ControlSetId string `json:"controlSetId,omitempty" yaml:"controlSetId,omitempty"`

	// Amazon Resource Name (ARN) of the IAM role.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
