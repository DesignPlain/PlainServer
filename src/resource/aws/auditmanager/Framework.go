package auditmanager

import types "DesignSphere_Server/src/resource/aws/types"

type Framework struct {
	// Description of the framework.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of the framework.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the framework. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Compliance type that the new custom framework supports, such as `CIS` or `HIPAA`.
	ComplianceType string `json:"complianceType,omitempty" yaml:"complianceType,omitempty"`

	/*
	   Control sets that are associated with the framework. See `control_sets` below.

	   The following arguments are optional:
	*/
	ControlSets []types.Auditmanager_FrameworkControlSet `json:"controlSets,omitempty" yaml:"controlSets,omitempty"`
}
