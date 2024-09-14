package securitycenter

import types "libds/gcp/types"

type SourceIamBinding struct {
	//
	Condition types.Securitycenter_SourceIamBindingCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	//
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`

	/*
	   The organization whose Cloud Security Command Center the Source
	   lives in.


	   - - -
	*/
	Organization string `json:"organization,omitempty" yaml:"organization,omitempty"`

	//
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	//
	Source string `json:"source,omitempty" yaml:"source,omitempty"`
}
