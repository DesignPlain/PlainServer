package dataform

type RepositoryIamPolicy struct {
	//
	PolicyData string `json:"policyData,omitempty" yaml:"policyData,omitempty"`

	//
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	//
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	//
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`
}
