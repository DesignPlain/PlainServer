package grafana

type RoleAssociation struct {
	// The AWS SSO group ids to be assigned the role given in `role`.
	GroupIds []string `json:"groupIds,omitempty" yaml:"groupIds,omitempty"`

	// The grafana role. Valid values can be found [here](https://docs.aws.amazon.com/grafana/latest/APIReference/API_UpdateInstruction.html#ManagedGrafana-Type-UpdateInstruction-role).
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	// The AWS SSO user ids to be assigned the role given in `role`.
	UserIds []string `json:"userIds,omitempty" yaml:"userIds,omitempty"`

	/*
	   The workspace id.

	   The following arguments are optional:
	*/
	WorkspaceId string `json:"workspaceId,omitempty" yaml:"workspaceId,omitempty"`
}
