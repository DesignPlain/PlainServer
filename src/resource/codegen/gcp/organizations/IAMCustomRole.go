package organizations

type IAMCustomRole struct {
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

	// The numeric ID of the organization in which you want to create a custom role.
	OrgId string `json:"orgId,omitempty" yaml:"orgId,omitempty"`

	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	Permissions []string `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	// The role id to use for this role.
	RoleId string `json:"roleId,omitempty" yaml:"roleId,omitempty"`
}
