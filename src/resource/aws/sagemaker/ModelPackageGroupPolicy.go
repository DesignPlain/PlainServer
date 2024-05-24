package sagemaker

type ModelPackageGroupPolicy struct {
	// The name of the model package group.
	ModelPackageGroupName string `json:"modelPackageGroupName,omitempty" yaml:"modelPackageGroupName,omitempty"`

	//
	ResourcePolicy string `json:"resourcePolicy,omitempty" yaml:"resourcePolicy,omitempty"`
}
