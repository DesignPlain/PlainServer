package types

type Vmwareengine_PrivateCloudManagementClusterNodeTypeConfig struct {
	/*
	   Customized number of cores available to each node of the type.
	   This number must always be one of `nodeType.availableCustomCoreCounts`.
	   If zero is provided max value from `nodeType.availableCustomCoreCounts` will be used.
	   This cannot be changed once the PrivateCloud is created.

	   - - -
	*/
	CustomCoreCount int `json:"customCoreCount,omitempty" yaml:"customCoreCount,omitempty"`

	// The number of nodes of this type in the cluster.
	NodeCount int `json:"nodeCount,omitempty" yaml:"nodeCount,omitempty"`

	// The identifier for this object. Format specified above.
	NodeTypeId string `json:"nodeTypeId,omitempty" yaml:"nodeTypeId,omitempty"`
}
