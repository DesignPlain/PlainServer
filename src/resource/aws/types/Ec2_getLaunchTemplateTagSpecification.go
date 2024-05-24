package types

type Ec2_getLaunchTemplateTagSpecification struct {
	//
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`

	// Map of tags, each pair of which must exactly match a pair on the desired Launch Template.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
