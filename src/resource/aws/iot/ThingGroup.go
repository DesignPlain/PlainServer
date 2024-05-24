package iot

import types "DesignSphere_Server/src/resource/aws/types"

type ThingGroup struct {
	// The name of the Thing Group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The name of the parent Thing Group.
	ParentGroupName string `json:"parentGroupName,omitempty" yaml:"parentGroupName,omitempty"`

	// The Thing Group properties. Defined below.
	Properties types.Iot_ThingGroupProperties `json:"properties,omitempty" yaml:"properties,omitempty"`

	// Key-value mapping of resource tags
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
