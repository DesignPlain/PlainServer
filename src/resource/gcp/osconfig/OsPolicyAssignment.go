package osconfig

import types "DesignSphere_Server/src/resource/gcp/types"

type OsPolicyAssignment struct {
	/*
	   OS policy assignment description. Length of the
	   description is limited to 1024 characters.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Filter to select VMs. Structure is
	   documented below.
	*/
	InstanceFilter types.Osconfig_OsPolicyAssignmentInstanceFilter `json:"instanceFilter,omitempty" yaml:"instanceFilter,omitempty"`

	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Resource name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   List of OS policies to be applied to the VMs.
	   Structure is documented below.
	*/
	OsPolicies []types.Osconfig_OsPolicyAssignmentOsPolicy `json:"osPolicies,omitempty" yaml:"osPolicies,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Rollout to deploy the OS policy assignment. A rollout
	   is triggered in the following situations: 1) OSPolicyAssignment is created.
	   2) OSPolicyAssignment is updated and the update contains changes to one of
	   the following fields: - instance_filter - os_policies 3) OSPolicyAssignment
	   is deleted. Structure is documented below.
	*/
	Rollout types.Osconfig_OsPolicyAssignmentRollout `json:"rollout,omitempty" yaml:"rollout,omitempty"`

	/*
	   Set to true to skip awaiting rollout
	   during resource creation and update.
	*/
	SkipAwaitRollout bool `json:"skipAwaitRollout,omitempty" yaml:"skipAwaitRollout,omitempty"`
}
