package types

type Container_getClusterNodePool struct {
	// Creates a unique name for the node pool beginning with the specified prefix. Conflicts with name.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// Specifies the configuration of queued provisioning
	QueuedProvisionings []Container_getClusterNodePoolQueuedProvisioning `json:"queuedProvisionings,omitempty" yaml:"queuedProvisionings,omitempty"`

	// The initial number of nodes for the pool. In regional or multi-zonal clusters, this is the number of nodes per zone. Changing this will force recreation of the resource.
	InitialNodeCount int `json:"initialNodeCount,omitempty" yaml:"initialNodeCount,omitempty"`

	// List of instance group URLs which have been assigned to this node pool.
	ManagedInstanceGroupUrls []string `json:"managedInstanceGroupUrls,omitempty" yaml:"managedInstanceGroupUrls,omitempty"`

	// The maximum number of pods per node in this node pool. Note that this does not work on node pools which are "route-based" - that is, node pools belonging to clusters that do not have IP Aliasing enabled.
	MaxPodsPerNode int `json:"maxPodsPerNode,omitempty" yaml:"maxPodsPerNode,omitempty"`

	// The name of the cluster.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The configuration of the nodepool
	NodeConfigs []Container_getClusterNodePoolNodeConfig `json:"nodeConfigs,omitempty" yaml:"nodeConfigs,omitempty"`

	// The resource URLs of the managed instance groups associated with this node pool.
	InstanceGroupUrls []string `json:"instanceGroupUrls,omitempty" yaml:"instanceGroupUrls,omitempty"`

	// Node management configuration, wherein auto-repair and auto-upgrade is configured.
	Managements []Container_getClusterNodePoolManagement `json:"managements,omitempty" yaml:"managements,omitempty"`

	// Networking configuration for this NodePool. If specified, it overrides the cluster-level defaults.
	NetworkConfigs []Container_getClusterNodePoolNetworkConfig `json:"networkConfigs,omitempty" yaml:"networkConfigs,omitempty"`

	// The number of nodes per instance group. This field can be used to update the number of nodes per instance group but should not be used alongside autoscaling.
	NodeCount int `json:"nodeCount,omitempty" yaml:"nodeCount,omitempty"`

	// The list of zones in which the node pool's nodes should be located. Nodes must be in the region of their regional cluster or in the same region as their cluster's zone for zonal clusters. If unspecified, the cluster-level node_locations will be used.
	NodeLocations []string `json:"nodeLocations,omitempty" yaml:"nodeLocations,omitempty"`

	// Specifies the node placement policy
	PlacementPolicies []Container_getClusterNodePoolPlacementPolicy `json:"placementPolicies,omitempty" yaml:"placementPolicies,omitempty"`

	// Specify node upgrade settings to change how many nodes GKE attempts to upgrade at once. The number of nodes upgraded simultaneously is the sum of max_surge and max_unavailable. The maximum number of nodes upgraded simultaneously is limited to 20.
	UpgradeSettings []Container_getClusterNodePoolUpgradeSetting `json:"upgradeSettings,omitempty" yaml:"upgradeSettings,omitempty"`

	//
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// Configuration required by cluster autoscaler to adjust the size of the node pool to the current cluster usage.
	Autoscalings []Container_getClusterNodePoolAutoscaling `json:"autoscalings,omitempty" yaml:"autoscalings,omitempty"`
}
