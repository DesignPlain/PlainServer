package dataproc

import types "DesignSphere_Server/src/resource/gcp/types"

type JobIAMBinding struct {
	//
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`

	/*
	   The project in which the job belongs. If it
	   is not provided, the provider will use a default.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The region in which the job belongs. If it
	   is not provided, the provider will use a default.
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   The role that should be applied. Only one
	   `gcp.dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
	   `[projects|organizations]/{parent-name}/roles/{role-name}`.

	   `gcp.dataproc.JobIAMPolicy` only:
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	//
	Condition types.Dataproc_JobIAMBindingCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	//
	JobId string `json:"jobId,omitempty" yaml:"jobId,omitempty"`
}
