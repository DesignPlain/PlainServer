package appstream

import types "libds/aws/types"

type Stack struct {
	// URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed. .
	FeedbackUrl string `json:"feedbackUrl,omitempty" yaml:"feedbackUrl,omitempty"`

	// URL that users are redirected to after their streaming session ends.
	RedirectUrl string `json:"redirectUrl,omitempty" yaml:"redirectUrl,omitempty"`

	/*
	   Configuration block for the storage connectors to enable.
	   See `storage_connectors` below.
	*/
	StorageConnectors []types.Appstream_StackStorageConnector `json:"storageConnectors,omitempty" yaml:"storageConnectors,omitempty"`

	/*
	   The streaming protocol you want your stack to prefer. This can be UDP or TCP. Currently, UDP is only supported in the Windows native client.
	   See `streaming_experience_settings` below.
	*/
	StreamingExperienceSettings types.Appstream_StackStreamingExperienceSettings `json:"streamingExperienceSettings,omitempty" yaml:"streamingExperienceSettings,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Set of configuration blocks defining the interface VPC endpoints. Users of the stack can connect to AppStream 2.0 only through the specified endpoints.
	   See `access_endpoints` below.
	*/
	AccessEndpoints []types.Appstream_StackAccessEndpoint `json:"accessEndpoints,omitempty" yaml:"accessEndpoints,omitempty"`

	/*
	   Settings for application settings persistence.
	   See `application_settings` below.
	*/
	ApplicationSettings types.Appstream_StackApplicationSettings `json:"applicationSettings,omitempty" yaml:"applicationSettings,omitempty"`

	// Domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions.
	EmbedHostDomains []string `json:"embedHostDomains,omitempty" yaml:"embedHostDomains,omitempty"`

	/*
	   Configuration block for the actions that are enabled or disabled for users during their streaming sessions. If not provided, these settings are configured automatically by AWS. If provided, the configuration should include a block for each configurable action.
	   See `user_settings` below.
	*/
	UserSettings []types.Appstream_StackUserSetting `json:"userSettings,omitempty" yaml:"userSettings,omitempty"`

	// Description for the AppStream stack.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Stack name to display.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Unique name for the AppStream stack.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
