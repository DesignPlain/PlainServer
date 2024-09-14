package quicksight

type GroupMembership struct {
	// The ID for the AWS account that the group is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId string `json:"awsAccountId,omitempty" yaml:"awsAccountId,omitempty"`

	// The name of the group in which the member will be added.
	GroupName string `json:"groupName,omitempty" yaml:"groupName,omitempty"`

	// The name of the member to add to the group.
	MemberName string `json:"memberName,omitempty" yaml:"memberName,omitempty"`

	// The namespace that you want the user to be a part of. Defaults to `default`.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}
