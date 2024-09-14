package grafana

type WorkspaceServiceAccountToken struct {
	// A name for the token to create. The name must be unique within the workspace.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Sets how long the token will be valid, in seconds. You can set the time up to 30 days in the future.
	SecondsToLive int `json:"secondsToLive,omitempty" yaml:"secondsToLive,omitempty"`

	// The ID of the service account for which to create a token.
	ServiceAccountId string `json:"serviceAccountId,omitempty" yaml:"serviceAccountId,omitempty"`

	// The Grafana workspace with which the service account token is associated.
	WorkspaceId string `json:"workspaceId,omitempty" yaml:"workspaceId,omitempty"`
}
