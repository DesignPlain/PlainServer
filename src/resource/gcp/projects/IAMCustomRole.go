package projects

type IAMCustomRole struct {
	// The camel case role id to use for this role. Cannot contain `-` characters.
	RoleId string `json:"roleId,omitempty" yaml:"roleId,omitempty"`

	/*
	   The current launch stage of the role.
	   Defaults to `GA`.
	   List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
	*/
	Stage string `json:"stage,omitempty" yaml:"stage,omitempty"`

	// A human-readable title for the role.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`

	// A human-readable description for the role.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	Permissions []string `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	/*
	   The project that the custom role will be created in.
	   Defaults to the provider project configuration.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
