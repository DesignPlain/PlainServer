package gkebackup

type RestorePlanIamPolicy struct {
	//
	PolicyData string `json:"policyData,omitempty" yaml:"policyData,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The region of the Restore Plan.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The full name of the BackupPlan Resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
