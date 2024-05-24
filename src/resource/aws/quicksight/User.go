package quicksight

type User struct {
	// The name of the IAM session to use when assuming roles that can embed QuickSight dashboards. Only valid for registering users using an assumed IAM role. Additionally, if registering multiple users using the same IAM role, each user needs to have a unique session name.
	SessionName string `json:"sessionName,omitempty" yaml:"sessionName,omitempty"`

	// The Amazon QuickSight user name that you want to create for the user you are registering. Only valid for registering a user with `identity_type` set to `QUICKSIGHT`.
	UserName string `json:"userName,omitempty" yaml:"userName,omitempty"`

	// The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
	UserRole string `json:"userRole,omitempty" yaml:"userRole,omitempty"`

	// The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId string `json:"awsAccountId,omitempty" yaml:"awsAccountId,omitempty"`

	// The email address of the user that you want to register.
	Email string `json:"email,omitempty" yaml:"email,omitempty"`

	// The ARN of the IAM user or role that you are registering with Amazon QuickSight.
	IamArn string `json:"iamArn,omitempty" yaml:"iamArn,omitempty"`

	// Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`. If `IAM` is specified, the `iam_arn` must also be specified.
	IdentityType string `json:"identityType,omitempty" yaml:"identityType,omitempty"`

	// The Amazon Quicksight namespace to create the user in. Defaults to `default`.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}
