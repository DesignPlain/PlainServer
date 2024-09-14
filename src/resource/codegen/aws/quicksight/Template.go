package quicksight

import types "libds/aws/types"

type Template struct {
	// Identifier for the template.
	TemplateId string `json:"templateId,omitempty" yaml:"templateId,omitempty"`

	/*
	   A description of the current template version being created/updated.

	   The following arguments are optional:
	*/
	VersionDescription string `json:"versionDescription,omitempty" yaml:"versionDescription,omitempty"`

	// AWS account ID.
	AwsAccountId string `json:"awsAccountId,omitempty" yaml:"awsAccountId,omitempty"`

	// Display name for the template.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A set of resource permissions on the template. Maximum of 64 items. See permissions.
	Permissions []types.Quicksight_TemplatePermission `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	// The entity that you are using as a source when you create the template (analysis or template). Only one of `definition` or `source_entity` should be configured. See source_entity.
	SourceEntity types.Quicksight_TemplateSourceEntity `json:"sourceEntity,omitempty" yaml:"sourceEntity,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
