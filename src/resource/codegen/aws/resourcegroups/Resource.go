package resourcegroups

type Resource struct {
	// The ARN of the resource to be added to the group.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`

	/*
	   The name or the ARN of the resource group to add resources to.

	   The following arguments are optional:
	*/
	GroupArn string `json:"groupArn,omitempty" yaml:"groupArn,omitempty"`
}
