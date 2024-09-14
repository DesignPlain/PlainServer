package types

type Osconfig_GuestPoliciesPackageRepository struct {
	/*
	   A Yum Repository.
	   Structure is documented below.
	*/
	Yum Osconfig_GuestPoliciesPackageRepositoryYum `json:"yum,omitempty" yaml:"yum,omitempty"`

	/*
	   A Zypper Repository.
	   Structure is documented below.
	*/
	Zypper Osconfig_GuestPoliciesPackageRepositoryZypper `json:"zypper,omitempty" yaml:"zypper,omitempty"`

	/*
	   An Apt Repository.
	   Structure is documented below.
	*/
	Apt Osconfig_GuestPoliciesPackageRepositoryApt `json:"apt,omitempty" yaml:"apt,omitempty"`

	/*
	   A Goo Repository.
	   Structure is documented below.
	*/
	Goo Osconfig_GuestPoliciesPackageRepositoryGoo `json:"goo,omitempty" yaml:"goo,omitempty"`
}
