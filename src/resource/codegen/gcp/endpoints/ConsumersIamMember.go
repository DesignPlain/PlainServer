package endpoints

import types "libds/gcp/types"

type ConsumersIamMember struct {
	//
	ServiceName string `json:"serviceName,omitempty" yaml:"serviceName,omitempty"`

	//
	Condition types.Endpoints_ConsumersIamMemberCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	//
	ConsumerProject string `json:"consumerProject,omitempty" yaml:"consumerProject,omitempty"`

	//
	Member string `json:"member,omitempty" yaml:"member,omitempty"`

	/*
	   The role that should be applied. Only one
	   `gcp.endpoints.ConsumersIamBinding` can be used per role. Note that custom roles must be of the format
	   `[projects|organizations]/{parent-name}/roles/{role-name}`.
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`
}
