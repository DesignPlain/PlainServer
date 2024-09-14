package ec2

type Tag struct {
	// The tag name.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// The ID of the EC2 resource to manage the tag for.
	ResourceId string `json:"resourceId,omitempty" yaml:"resourceId,omitempty"`

	// The value of the tag.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
