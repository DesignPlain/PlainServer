package dataproc

import types "DesignSphere_Server/src/resource/gcp/types"

type AutoscalingPolicy struct {
	/*
	   Describes how the autoscaler will operate for primary workers.
	   Structure is documented below.
	*/
	WorkerConfig types.Dataproc_AutoscalingPolicyWorkerConfig `json:"workerConfig,omitempty" yaml:"workerConfig,omitempty"`

	/*
	   Basic algorithm for autoscaling.
	   Structure is documented below.
	*/
	BasicAlgorithm types.Dataproc_AutoscalingPolicyBasicAlgorithm `json:"basicAlgorithm,omitempty" yaml:"basicAlgorithm,omitempty"`

	/*
	   The  location where the autoscaling policy should reside.
	   The default value is `global`.
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	   and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
	   3 and 50 characters.


	   - - -
	*/
	PolicyId string `json:"policyId,omitempty" yaml:"policyId,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Describes how the autoscaler will operate for secondary workers.
	   Structure is documented below.
	*/
	SecondaryWorkerConfig types.Dataproc_AutoscalingPolicySecondaryWorkerConfig `json:"secondaryWorkerConfig,omitempty" yaml:"secondaryWorkerConfig,omitempty"`
}
