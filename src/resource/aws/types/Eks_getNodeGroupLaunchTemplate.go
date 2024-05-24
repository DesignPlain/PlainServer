package types

type Eks_getNodeGroupLaunchTemplate struct {
	// The ID of the launch template.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Name of the AutoScaling Group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Kubernetes version.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
