package types

type Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourcePkgDeb struct {
	/*
	   Whether dependencies should also be installed. -
	   install when false: `dpkg -i package` - install when true: `apt-get update
	   && apt-get -y install package.deb`
	*/
	PullDeps bool `json:"pullDeps,omitempty" yaml:"pullDeps,omitempty"`

	/*
	   A deb package. Structure is
	   documented below.
	*/
	Source Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourcePkgDebSource `json:"source,omitempty" yaml:"source,omitempty"`
}
