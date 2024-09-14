package backup

import types "libds/aws/types"

type Selection struct {
	// An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to exclude from a backup plan.
	NotResources []string `json:"notResources,omitempty" yaml:"notResources,omitempty"`

	// The backup plan ID to be associated with the selection of resources.
	PlanId string `json:"planId,omitempty" yaml:"planId,omitempty"`

	// An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan.
	Resources []string `json:"resources,omitempty" yaml:"resources,omitempty"`

	// Tag-based conditions used to specify a set of resources to assign to a backup plan.
	SelectionTags []types.Backup_SelectionSelectionTag `json:"selectionTags,omitempty" yaml:"selectionTags,omitempty"`

	// A list of conditions that you define to assign resources to your backup plans using tags.
	Conditions []types.Backup_SelectionCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`

	// The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/access-control.html#managed-policies) for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
	IamRoleArn string `json:"iamRoleArn,omitempty" yaml:"iamRoleArn,omitempty"`

	// The display name of a resource selection document.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
