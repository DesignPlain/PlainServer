package types

type Quicksight_DashboardPermission struct {
	// ARN of the principal. See the [ResourcePermission documentation](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_ResourcePermission.html) for the applicable ARN values.
	Principal string `json:"principal,omitempty" yaml:"principal,omitempty"`

	// List of IAM actions to grant or revoke permissions on.
	Actions []string `json:"actions,omitempty" yaml:"actions,omitempty"`
}
