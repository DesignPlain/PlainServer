package quicksight

import types "libds/aws/types"

type Analysis struct {
	// AWS account ID.
	AwsAccountId string `json:"awsAccountId,omitempty" yaml:"awsAccountId,omitempty"`

	// The entity that you are using as a source when you create the analysis (template). Only one of `definition` or `source_entity` should be configured. See source_entity.
	SourceEntity types.Quicksight_AnalysisSourceEntity `json:"sourceEntity,omitempty" yaml:"sourceEntity,omitempty"`

	// The Amazon Resource Name (ARN) of the theme that is being used for this analysis. The theme ARN must exist in the same AWS account where you create the analysis.
	ThemeArn string `json:"themeArn,omitempty" yaml:"themeArn,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Identifier for the analysis.
	AnalysisId string `json:"analysisId,omitempty" yaml:"analysisId,omitempty"`

	/*
	   Display name for the analysis.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The parameters for the creation of the analysis, which you want to use to override the default settings. An analysis can have any type of parameters, and some parameters might accept multiple values. See parameters.
	Parameters types.Quicksight_AnalysisParameters `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// A set of resource permissions on the analysis. Maximum of 64 items. See permissions.
	Permissions []types.Quicksight_AnalysisPermission `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	// A value that specifies the number of days that Amazon QuickSight waits before it deletes the analysis. Use `0` to force deletion without recovery. Minimum value of `7`. Maximum value of `30`. Default to `30`.
	RecoveryWindowInDays int `json:"recoveryWindowInDays,omitempty" yaml:"recoveryWindowInDays,omitempty"`
}
