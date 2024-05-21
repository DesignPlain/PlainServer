package types

type Dataproc_ClusterClusterConfigAuxiliaryNodeGroupNodeGroup struct {
	/*
	   The name of the cluster, unique within the project and
	   zone.

	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The node group instance group configuration.
	NodeGroupConfig Dataproc_ClusterClusterConfigAuxiliaryNodeGroupNodeGroupNodeGroupConfig `json:"nodeGroupConfig,omitempty" yaml:"nodeGroupConfig,omitempty"`

	/*
	   Node group roles.
	   One of `"DRIVER"`.
	*/
	Roles []string `json:"roles,omitempty" yaml:"roles,omitempty"`
}
