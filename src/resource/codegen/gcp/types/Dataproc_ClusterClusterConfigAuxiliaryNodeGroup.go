package types

type Dataproc_ClusterClusterConfigAuxiliaryNodeGroup struct {
	// A node group ID. Generated if not specified. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of from 3 to 33 characters.
	NodeGroupId string `json:"nodeGroupId,omitempty" yaml:"nodeGroupId,omitempty"`

	// Node group configuration.
	NodeGroups []Dataproc_ClusterClusterConfigAuxiliaryNodeGroupNodeGroup `json:"nodeGroups,omitempty" yaml:"nodeGroups,omitempty"`
}
