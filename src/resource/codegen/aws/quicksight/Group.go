package quicksight

type Group struct {
	// The namespace. Currently, you should set this to `default`.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	// The ID for the AWS account that the group is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId string `json:"awsAccountId,omitempty" yaml:"awsAccountId,omitempty"`

	// A description for the group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A name for the group.
	GroupName string `json:"groupName,omitempty" yaml:"groupName,omitempty"`
}
