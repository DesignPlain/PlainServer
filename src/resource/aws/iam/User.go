package iam

type User struct {
	// Key-value mapping of tags for the IAM user. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   When destroying this user, destroy even if it
	   has non-provider-managed IAM access keys, login profile or MFA devices. Without `force_destroy`
	   a user with non-provider-managed access keys and login profile will fail to be destroyed.
	*/
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	// The user's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. User names are not distinguished by case. For example, you cannot create users named both "TESTUSER" and "testuser".
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Path in which to create the user.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// The ARN of the policy that is used to set the permissions boundary for the user.
	PermissionsBoundary string `json:"permissionsBoundary,omitempty" yaml:"permissionsBoundary,omitempty"`
}
