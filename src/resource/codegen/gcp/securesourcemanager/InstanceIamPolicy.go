package securesourcemanager

type InstanceIamPolicy struct {
	//
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	//
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	//
	PolicyData string `json:"policyData,omitempty" yaml:"policyData,omitempty"`

	//
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
