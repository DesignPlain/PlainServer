package grafana

type WorkspaceApiKey struct {
	// Specifies the name of the API key. Key names must be unique to the workspace.
	KeyName string `json:"keyName,omitempty" yaml:"keyName,omitempty"`

	// Specifies the permission level of the API key. Valid values are `VIEWER`, `EDITOR`, or `ADMIN`.
	KeyRole string `json:"keyRole,omitempty" yaml:"keyRole,omitempty"`

	// Specifies the time in seconds until the API key expires. Keys can be valid for up to 30 days.
	SecondsToLive int `json:"secondsToLive,omitempty" yaml:"secondsToLive,omitempty"`

	// The ID of the workspace that the API key is valid for.
	WorkspaceId string `json:"workspaceId,omitempty" yaml:"workspaceId,omitempty"`
}
