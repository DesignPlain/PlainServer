package endpoints

import types "DesignSphere_Server/src/resource/gcp/types"

type ServiceIamBinding struct {
	/*
	   The role that should be applied. Only one
	   `gcp.endpoints.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
	   `[projects|organizations]/{parent-name}/roles/{role-name}`.
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	//
	ServiceName string `json:"serviceName,omitempty" yaml:"serviceName,omitempty"`

	//
	Condition types.Endpoints_ServiceIamBindingCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	//
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`
}
