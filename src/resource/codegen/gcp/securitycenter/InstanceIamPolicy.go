package securitycenter

type InstanceIamPolicy struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The region of the Data Fusion instance.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The ID of the instance or a fully qualified identifier for the instance.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	PolicyData string `json:"policyData,omitempty" yaml:"policyData,omitempty"`
}
