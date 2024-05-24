package connect

import types "DesignSphere_Server/src/resource/aws/types"

type UserHierarchyStructure struct {
	// A block that defines the hierarchy structure's levels. The `hierarchy_structure` block is documented below.
	HierarchyStructure types.Connect_UserHierarchyStructureHierarchyStructure `json:"hierarchyStructure,omitempty" yaml:"hierarchyStructure,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`
}
