package osconfig

import types "DesignSphere_Server/src/resource/gcp/types"

type GuestPolicies struct {
	// Description of the guest policy. Length of the description is limited to 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The etag for this guest policy. If this is provided on update, it must match the server's etag.
	Etag string `json:"etag,omitempty" yaml:"etag,omitempty"`

	/*
	   The logical name of the guest policy in the project with the following restrictions:
	   - Must contain only lowercase letters, numbers, and hyphens.
	   - Must start with a letter.
	   - Must be between 1-63 characters.
	   - Must end with a number or a letter.
	   - Must be unique within the project.
	*/
	GuestPolicyId string `json:"guestPolicyId,omitempty" yaml:"guestPolicyId,omitempty"`

	/*
	   A list of package repositories to configure on the VM instance.
	   This is done before any other configs are applied so they can use these repos.
	   Package repositories are only configured if the corresponding package manager(s) are available.
	   Structure is documented below.
	*/
	PackageRepositories []types.Osconfig_GuestPoliciesPackageRepository `json:"packageRepositories,omitempty" yaml:"packageRepositories,omitempty"`

	/*
	   The software packages to be managed by this policy.
	   Structure is documented below.
	*/
	Packages []types.Osconfig_GuestPoliciesPackage `json:"packages,omitempty" yaml:"packages,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   A list of Recipes to install on the VM instance.
	   Structure is documented below.
	*/
	Recipes []types.Osconfig_GuestPoliciesRecipe `json:"recipes,omitempty" yaml:"recipes,omitempty"`

	/*
	   Specifies the VM instances that are assigned to this policy. This allows you to target sets
	   or groups of VM instances by different parameters such as labels, names, OS, or zones.
	   If left empty, all VM instances underneath this policy are targeted.
	   At the same level in the resource hierarchy (that is within a project), the service prevents
	   the creation of multiple policies that conflict with each other.
	   For more information, see how the service
	   [handles assignment conflicts](https://cloud.google.com/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
	   Structure is documented below.
	*/
	Assignment types.Osconfig_GuestPoliciesAssignment `json:"assignment,omitempty" yaml:"assignment,omitempty"`
}
