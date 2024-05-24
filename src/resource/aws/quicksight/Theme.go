package quicksight

import types "DesignSphere_Server/src/resource/aws/types"

type Theme struct {
	// AWS account ID.
	AwsAccountId string `json:"awsAccountId,omitempty" yaml:"awsAccountId,omitempty"`

	// The ID of the theme that a custom theme will inherit from. All themes inherit from one of the starting themes defined by Amazon QuickSight. For a list of the starting themes, use ListThemes or choose Themes from within an analysis.
	BaseThemeId string `json:"baseThemeId,omitempty" yaml:"baseThemeId,omitempty"`

	/*
	   The theme configuration, which contains the theme display properties. See configuration.

	   The following arguments are optional:
	*/
	Configuration types.Quicksight_ThemeConfiguration `json:"configuration,omitempty" yaml:"configuration,omitempty"`

	// Display name of the theme.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A set of resource permissions on the theme. Maximum of 64 items. See permissions.
	Permissions []types.Quicksight_ThemePermission `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Identifier of the theme.
	ThemeId string `json:"themeId,omitempty" yaml:"themeId,omitempty"`

	// A description of the current theme version being created/updated.
	VersionDescription string `json:"versionDescription,omitempty" yaml:"versionDescription,omitempty"`
}
