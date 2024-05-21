package types

type Osconfig_GuestPoliciesAssignmentOsType struct {
	/*
	   Targets VM instances with OS Inventory enabled and having the following OS architecture.

	   - - -
	*/
	OsArchitecture string `json:"osArchitecture,omitempty" yaml:"osArchitecture,omitempty"`

	// Targets VM instances with OS Inventory enabled and having the following OS short name, for example "debian" or "windows".
	OsShortName string `json:"osShortName,omitempty" yaml:"osShortName,omitempty"`

	// Targets VM instances with OS Inventory enabled and having the following following OS version.
	OsVersion string `json:"osVersion,omitempty" yaml:"osVersion,omitempty"`
}
