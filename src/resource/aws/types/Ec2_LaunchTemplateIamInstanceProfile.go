package types

type Ec2_LaunchTemplateIamInstanceProfile struct {
	// The Amazon Resource Name (ARN) of the instance profile. Conflicts with `name`.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// The name of the instance profile.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
