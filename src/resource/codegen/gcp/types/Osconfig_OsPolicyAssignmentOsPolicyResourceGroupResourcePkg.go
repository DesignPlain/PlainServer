package types

type Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourcePkg struct {
	/*
	   An rpm package file. Structure is
	   documented below.
	*/
	Rpm Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourcePkgRpm `json:"rpm,omitempty" yaml:"rpm,omitempty"`

	/*
	   A package managed by YUM. Structure is
	   documented below.
	*/
	Yum Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourcePkgYum `json:"yum,omitempty" yaml:"yum,omitempty"`

	/*
	   A package managed by Zypper. Structure is
	   documented below.
	*/
	Zypper Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourcePkgZypper `json:"zypper,omitempty" yaml:"zypper,omitempty"`

	/*
	   A package managed by Apt. Structure is
	   documented below.
	*/
	Apt Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourcePkgApt `json:"apt,omitempty" yaml:"apt,omitempty"`

	/*
	   A deb package file. Structure is
	   documented below.
	*/
	Deb Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourcePkgDeb `json:"deb,omitempty" yaml:"deb,omitempty"`

	/*
	   The desired state the agent should maintain for
	   this package. Possible values are: `DESIRED_STATE_UNSPECIFIED`, `INSTALLED`,
	   `REMOVED`.
	*/
	DesiredState string `json:"desiredState,omitempty" yaml:"desiredState,omitempty"`

	/*
	   A package managed by GooGet. Structure is
	   documented below.
	*/
	Googet Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourcePkgGooget `json:"googet,omitempty" yaml:"googet,omitempty"`

	/*
	   An MSI package. Structure is
	   documented below.
	*/
	Msi Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourcePkgMsi `json:"msi,omitempty" yaml:"msi,omitempty"`
}
