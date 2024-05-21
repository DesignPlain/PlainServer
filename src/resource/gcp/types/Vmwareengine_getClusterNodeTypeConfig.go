package types

type Vmwareengine_getClusterNodeTypeConfig struct {
	// The number of nodes of this type in the cluster.
	NodeCount int `json:"nodeCount,omitempty" yaml:"nodeCount,omitempty"`

	//
	NodeTypeId string `json:"nodeTypeId,omitempty" yaml:"nodeTypeId,omitempty"`

	/*
	   Customized number of cores available to each node of the type.
	   This number must always be one of 'nodeType.availableCustomCoreCounts'.
	   If zero is provided max value from 'nodeType.availableCustomCoreCounts' will be used.
	   Once the customer is created then corecount cannot be changed.
	*/
	CustomCoreCount int `json:"customCoreCount,omitempty" yaml:"customCoreCount,omitempty"`
}
