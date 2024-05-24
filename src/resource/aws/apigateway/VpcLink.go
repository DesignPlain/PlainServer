package apigateway

type VpcLink struct {
	// List of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	TargetArn string `json:"targetArn,omitempty" yaml:"targetArn,omitempty"`

	// Description of the VPC link.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name used to label and identify the VPC link.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
