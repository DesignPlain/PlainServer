package vertex

import types "DesignSphere_Server/src/resource/gcp/types"

type AiEndpointIamMember struct {
	//
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	//
	Member string `json:"member,omitempty" yaml:"member,omitempty"`

	//
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	//
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	//
	Condition types.Vertex_AiEndpointIamMemberCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	//
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
}
