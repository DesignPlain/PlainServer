package types

type Osconfig_OsPolicyAssignmentOsPolicyResourceGroupInventoryFilter struct {
	// The OS short name
	OsShortName string `json:"osShortName,omitempty" yaml:"osShortName,omitempty"`

	/*
	   The OS version Prefix matches are supported if
	   asterisk(-) is provided as the last character. For example, to match all
	   versions with a major version of `7`, specify the following value for this
	   field `7.-` An empty string matches all OS versions.
	*/
	OsVersion string `json:"osVersion,omitempty" yaml:"osVersion,omitempty"`
}
