package types

type Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourcePkgMsi struct {
	/*
	   Additional properties to use during installation.
	   This should be in the format of Property=Setting. Appended to the defaults
	   of `ACTION=INSTALL REBOOT=ReallySuppress`.
	*/
	Properties []string `json:"properties,omitempty" yaml:"properties,omitempty"`

	/*
	   The MSI package. Structure is
	   documented below.
	*/
	Source Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourcePkgMsiSource `json:"source,omitempty" yaml:"source,omitempty"`
}
