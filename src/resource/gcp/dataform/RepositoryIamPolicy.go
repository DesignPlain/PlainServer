package dataform

type RepositoryIamPolicy struct {
	//
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	//
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`

	//
	PolicyData string `json:"policyData,omitempty" yaml:"policyData,omitempty"`

	//
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
