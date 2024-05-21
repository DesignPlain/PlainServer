package clouddeploy

type TargetIamPolicy struct {
	//
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	PolicyData string `json:"policyData,omitempty" yaml:"policyData,omitempty"`

	//
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	//
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}
