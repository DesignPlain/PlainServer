package gkebackup

import types "DesignSphere_Server/src/resource/gcp/types"

type RestorePlanIamBinding struct {
	//
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`

	// The full name of the BackupPlan Resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	//
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	//
	Condition types.Gkebackup_RestorePlanIamBindingCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	// The region of the Restore Plan.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}
