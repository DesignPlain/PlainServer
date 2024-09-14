package sagemaker

type ModelPackageGroup struct {
	// A description for the model group.
	ModelPackageGroupDescription string `json:"modelPackageGroupDescription,omitempty" yaml:"modelPackageGroupDescription,omitempty"`

	// The name of the model group.
	ModelPackageGroupName string `json:"modelPackageGroupName,omitempty" yaml:"modelPackageGroupName,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
