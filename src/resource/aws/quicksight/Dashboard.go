package quicksight

import types "DesignSphere_Server/src/resource/aws/types"

type Dashboard struct {
	// Options for publishing the dashboard. See dashboard_publish_options.
	DashboardPublishOptions types.Quicksight_DashboardDashboardPublishOptions `json:"dashboardPublishOptions,omitempty" yaml:"dashboardPublishOptions,omitempty"`

	// Display name for the dashboard.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A set of resource permissions on the dashboard. Maximum of 64 items. See permissions.
	Permissions []types.Quicksight_DashboardPermission `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	// The entity that you are using as a source when you create the dashboard (template). Only one of `definition` or `source_entity` should be configured. See source_entity.
	SourceEntity types.Quicksight_DashboardSourceEntity `json:"sourceEntity,omitempty" yaml:"sourceEntity,omitempty"`

	// The Amazon Resource Name (ARN) of the theme that is being used for this dashboard. The theme ARN must exist in the same AWS account where you create the dashboard.
	ThemeArn string `json:"themeArn,omitempty" yaml:"themeArn,omitempty"`

	/*
	   A description of the current dashboard version being created/updated.

	   The following arguments are optional:
	*/
	VersionDescription string `json:"versionDescription,omitempty" yaml:"versionDescription,omitempty"`

	// AWS account ID.
	AwsAccountId string `json:"awsAccountId,omitempty" yaml:"awsAccountId,omitempty"`

	// Identifier for the dashboard.
	DashboardId string `json:"dashboardId,omitempty" yaml:"dashboardId,omitempty"`

	// The parameters for the creation of the dashboard, which you want to use to override the default settings. A dashboard can have any type of parameters, and some parameters might accept multiple values. See parameters.
	Parameters types.Quicksight_DashboardParameters `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
