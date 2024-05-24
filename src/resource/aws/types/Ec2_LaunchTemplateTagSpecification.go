package types

type Ec2_LaunchTemplateTagSpecification struct {
	// A map of tags to assign to the resource.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The type of resource to tag.
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`
}
