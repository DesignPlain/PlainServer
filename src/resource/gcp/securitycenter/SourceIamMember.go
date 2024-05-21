package securitycenter

import types "DesignSphere_Server/src/resource/gcp/types"

type SourceIamMember struct {
	//
	Member string `json:"member,omitempty" yaml:"member,omitempty"`

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

	//
	Condition types.Securitycenter_SourceIamMemberCondition `json:"condition,omitempty" yaml:"condition,omitempty"`
}
