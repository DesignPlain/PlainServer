package grafana

type WorkspaceServiceAccount struct {
	// The permission level to use for this service account. For more information about the roles and the permissions each has, see the [User roles](https://docs.aws.amazon.com/grafana/latest/userguide/Grafana-user-roles.html) documentation.
	GrafanaRole string `json:"grafanaRole,omitempty" yaml:"grafanaRole,omitempty"`

	// A name for the service account. The name must be unique within the workspace, as it determines the ID associated with the service account.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The Grafana workspace with which the service account is associated.
	WorkspaceId string `json:"workspaceId,omitempty" yaml:"workspaceId,omitempty"`
}
