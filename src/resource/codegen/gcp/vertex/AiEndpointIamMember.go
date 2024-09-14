package vertex

import types "libds/gcp/types"

type AiEndpointIamMember struct {
	//
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	//
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	//
	Condition types.Vertex_AiEndpointIamMemberCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	//
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	//
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	//
	Member string `json:"member,omitempty" yaml:"member,omitempty"`
}
