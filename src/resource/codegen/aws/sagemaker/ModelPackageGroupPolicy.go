package sagemaker

type ModelPackageGroupPolicy struct {
	//
	ResourcePolicy string `json:"resourcePolicy,omitempty" yaml:"resourcePolicy,omitempty"`

	// The name of the model package group.
	ModelPackageGroupName string `json:"modelPackageGroupName,omitempty" yaml:"modelPackageGroupName,omitempty"`
}
