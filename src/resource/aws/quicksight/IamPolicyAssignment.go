package quicksight

import types "DesignSphere_Server/src/resource/aws/types"

type IamPolicyAssignment struct {
	// Name of the assignment.
	AssignmentName string `json:"assignmentName,omitempty" yaml:"assignmentName,omitempty"`

	/*
	   Status of the assignment. Valid values are `ENABLED`, `DISABLED`, and `DRAFT`.

	   The following arguments are optional:
	*/
	AssignmentStatus string `json:"assignmentStatus,omitempty" yaml:"assignmentStatus,omitempty"`

	// AWS account ID.
	AwsAccountId string `json:"awsAccountId,omitempty" yaml:"awsAccountId,omitempty"`

	// Amazon QuickSight users, groups, or both to assign the policy to. See `identities` block.
	Identities types.Quicksight_IamPolicyAssignmentIdentities `json:"identities,omitempty" yaml:"identities,omitempty"`

	// Namespace that contains the assignment. Defaults to `default`.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	// ARN of the IAM policy to apply to the Amazon QuickSight users and groups specified in this assignment.
	PolicyArn string `json:"policyArn,omitempty" yaml:"policyArn,omitempty"`
}
