package codebuild

import types "libds/aws/types"

type Webhook struct {
	// A regular expression used to determine which branches get built. Default is all branches are built. We recommend using `filter_group` over `branch_filter`.
	BranchFilter string `json:"branchFilter,omitempty" yaml:"branchFilter,omitempty"`

	// The type of build this webhook will trigger. Valid values for this parameter are: `BUILD`, `BUILD_BATCH`.
	BuildType string `json:"buildType,omitempty" yaml:"buildType,omitempty"`

	// Information about the webhook's trigger. Filter group blocks are documented below.
	FilterGroups []types.Codebuild_WebhookFilterGroup `json:"filterGroups,omitempty" yaml:"filterGroups,omitempty"`

	// The name of the build project.
	ProjectName string `json:"projectName,omitempty" yaml:"projectName,omitempty"`

	// Scope configuration for global or organization webhooks. Scope configuration blocks are documented below.
	ScopeConfiguration types.Codebuild_WebhookScopeConfiguration `json:"scopeConfiguration,omitempty" yaml:"scopeConfiguration,omitempty"`
}
