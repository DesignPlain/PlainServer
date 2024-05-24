package iot

import types "DesignSphere_Server/src/resource/aws/types"

type ProvisioningTemplate struct {
	// Creates a pre-provisioning hook template. Details below.
	PreProvisioningHook types.Iot_ProvisioningTemplatePreProvisioningHook `json:"preProvisioningHook,omitempty" yaml:"preProvisioningHook,omitempty"`

	// The role ARN for the role associated with the fleet provisioning template. This IoT role grants permission to provision a device.
	ProvisioningRoleArn string `json:"provisioningRoleArn,omitempty" yaml:"provisioningRoleArn,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The JSON formatted contents of the fleet provisioning template.
	TemplateBody string `json:"templateBody,omitempty" yaml:"templateBody,omitempty"`

	// The type you define in a provisioning template.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The description of the fleet provisioning template.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// True to enable the fleet provisioning template, otherwise false.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The name of the fleet provisioning template.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
