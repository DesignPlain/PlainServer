package rbin

import types "DesignSphere_Server/src/resource/aws/types"

type Rule struct {
	// The retention rule description.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Information about the retention rule lock configuration. See `lock_configuration` below.
	LockConfiguration types.Rbin_RuleLockConfiguration `json:"lockConfiguration,omitempty" yaml:"lockConfiguration,omitempty"`

	// Specifies the resource tags to use to identify resources that are to be retained by a tag-level retention rule. See `resource_tags` below.
	ResourceTags []types.Rbin_RuleResourceTag `json:"resourceTags,omitempty" yaml:"resourceTags,omitempty"`

	// The resource type to be retained by the retention rule. Valid values are `EBS_SNAPSHOT` and `EC2_IMAGE`.
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`

	/*
	   Information about the retention period for which the retention rule is to retain resources. See `retention_period` below.

	   The following arguments are optional:
	*/
	RetentionPeriod types.Rbin_RuleRetentionPeriod `json:"retentionPeriod,omitempty" yaml:"retentionPeriod,omitempty"`

	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
