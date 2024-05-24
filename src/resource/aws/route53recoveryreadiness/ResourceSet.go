package route53recoveryreadiness

import types "DesignSphere_Server/src/resource/aws/types"

type ResourceSet struct {
	// Type of the resources in the resource set.
	ResourceSetType string `json:"resourceSetType,omitempty" yaml:"resourceSetType,omitempty"`

	/*
	   List of resources to add to this resource set. See below.

	   The following arguments are optional:
	*/
	Resources []types.Route53recoveryreadiness_ResourceSetResource `json:"resources,omitempty" yaml:"resources,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Unique name describing the resource set.
	ResourceSetName string `json:"resourceSetName,omitempty" yaml:"resourceSetName,omitempty"`
}
