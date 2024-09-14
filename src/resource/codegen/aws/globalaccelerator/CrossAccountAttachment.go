package globalaccelerator

import types "libds/aws/types"

type CrossAccountAttachment struct {
	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Name of the Cross Account Attachment.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// List of AWS account IDs that are allowed to associate resources with the accelerator.
	Principals []string `json:"principals,omitempty" yaml:"principals,omitempty"`

	// List of resources to be associated with the accelerator.
	Resources []types.Globalaccelerator_CrossAccountAttachmentResource `json:"resources,omitempty" yaml:"resources,omitempty"`
}
