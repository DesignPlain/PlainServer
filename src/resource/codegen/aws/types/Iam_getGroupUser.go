package types

type Iam_getGroupUser struct {
	// Stable and unique string identifying the IAM user.
	UserId string `json:"userId,omitempty" yaml:"userId,omitempty"`

	// Name of the IAM user.
	UserName string `json:"userName,omitempty" yaml:"userName,omitempty"`

	// User ARN.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// Path to the IAM user.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}
