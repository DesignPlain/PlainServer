package types

type Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourcePkgRpm struct {
	/*
	   An rpm package. Structure is
	   documented below.
	*/
	Source Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourcePkgRpmSource `json:"source,omitempty" yaml:"source,omitempty"`

	/*
	   Whether dependencies should also be installed. -
	   install when false: `rpm --upgrade --replacepkgs package.rpm` - install when
	   true: `yum -y install package.rpm` or `zypper -y install package.rpm`
	*/
	PullDeps bool `json:"pullDeps,omitempty" yaml:"pullDeps,omitempty"`
}
