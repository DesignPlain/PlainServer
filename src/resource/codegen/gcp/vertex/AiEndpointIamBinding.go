package vertex

import types "libds/gcp/types"

type AiEndpointIamBinding struct {
	//
	Condition types.Vertex_AiEndpointIamBindingCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	//
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	//
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	//
	Members []string `json:"members,omitempty" yaml:"members,omitempty"`

	//
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	//
	Role string `json:"role,omitempty" yaml:"role,omitempty"`
}
