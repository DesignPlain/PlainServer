package types

type Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResource struct {
	/*
	   File resource Structure is
	   documented below.
	*/
	File Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourceFile `json:"file,omitempty" yaml:"file,omitempty"`

	/*
	   The id of the resource with the following restrictions:

	   -   Must contain only lowercase letters, numbers, and hyphens.
	   -   Must start with a letter.
	   -   Must be between 1-63 characters.
	   -   Must end with a number or a letter.
	   -   Must be unique within the OS policy.
	*/
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	/*
	   Package resource Structure is
	   documented below.
	*/
	Pkg Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourcePkg `json:"pkg,omitempty" yaml:"pkg,omitempty"`

	/*
	   Package repository resource Structure is
	   documented below.
	*/
	Repository Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourceRepository `json:"repository,omitempty" yaml:"repository,omitempty"`

	/*
	   Exec resource Structure is
	   documented below.
	*/
	Exec Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourceExec `json:"exec,omitempty" yaml:"exec,omitempty"`
}
