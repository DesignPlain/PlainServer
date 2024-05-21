package types

type Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourceRepository struct {
	/*
	   An Apt Repository. Structure is
	   documented below.
	*/
	Apt Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourceRepositoryApt `json:"apt,omitempty" yaml:"apt,omitempty"`

	/*
	   A Goo Repository. Structure is
	   documented below.
	*/
	Goo Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourceRepositoryGoo `json:"goo,omitempty" yaml:"goo,omitempty"`

	/*
	   A Yum Repository. Structure is
	   documented below.
	*/
	Yum Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourceRepositoryYum `json:"yum,omitempty" yaml:"yum,omitempty"`

	/*
	   A Zypper Repository. Structure is
	   documented below.
	*/
	Zypper Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourceRepositoryZypper `json:"zypper,omitempty" yaml:"zypper,omitempty"`
}
