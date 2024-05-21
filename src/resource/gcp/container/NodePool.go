package container

import types "DesignSphere_Server/src/resource/gcp/types"

type NodePool struct {
	/*
	   The cluster to create the node pool for. Cluster must be present in `location` provided for clusters. May be specified in the format `projects/{{project}}/locations/{{location}}/clusters/{{cluster}}` or as just the name of the cluster.

	   - - -
	*/
	Cluster string `json:"cluster,omitempty" yaml:"cluster,omitempty"`

	/*
	   Node management configuration, wherein auto-repair and
	   auto-upgrade is configured. Structure is documented below.
	*/
	Management types.Container_NodePoolManagement `json:"management,omitempty" yaml:"management,omitempty"`

	/*
	   The maximum number of pods per node in this node pool.
	   Note that this does not work on node pools which are "route-based" - that is, node
	   pools belonging to clusters that do not have IP Aliasing enabled.
	   See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
	   for more information.
	*/
	MaxPodsPerNode int `json:"maxPodsPerNode,omitempty" yaml:"maxPodsPerNode,omitempty"`

	/*
	   The network configuration of the pool. Such as
	   configuration for [Adding Pod IP address ranges](https://cloud.google.com/kubernetes-engine/docs/how-to/multi-pod-cidr)) to the node pool. Or enabling private nodes. Structure is
	   documented below
	*/
	NetworkConfig types.Container_NodePoolNetworkConfig `json:"networkConfig,omitempty" yaml:"networkConfig,omitempty"`

	/*
	   The number of nodes per instance group. This field can be used to
	   update the number of nodes per instance group but should not be used alongside `autoscaling`.
	*/
	NodeCount int `json:"nodeCount,omitempty" yaml:"nodeCount,omitempty"`

	/*
	   Creates a unique name for the node pool beginning
	   with the specified prefix. Conflicts with `name`.
	*/
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	/*
	   The ID of the project in which to create the node pool. If blank,
	   the provider-configured project will be used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Specifies node pool-level settings of queued provisioning.
	   Structure is documented below.

	   <a name="nested_autoscaling"></a>The `autoscaling` block supports (either total or per zone limits are required):
	*/
	QueuedProvisioning types.Container_NodePoolQueuedProvisioning `json:"queuedProvisioning,omitempty" yaml:"queuedProvisioning,omitempty"`

	/*
	   The initial number of nodes for the pool. In
	   regional or multi-zonal clusters, this is the number of nodes per zone. Changing
	   this will force recreation of the resource. WARNING: Resizing your node pool manually
	   may change this value in your existing cluster, which will trigger destruction
	   and recreation on the next provider run (to rectify the discrepancy).  If you don't
	   need this value, don't set it.  If you do need it, you can use a lifecycle block to
	   ignore subsqeuent changes to this field.
	*/
	InitialNodeCount int `json:"initialNodeCount,omitempty" yaml:"initialNodeCount,omitempty"`

	/*
	   The location (region or zone) of the cluster.

	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The list of zones in which the node pool's nodes should be located. Nodes must
	   be in the region of their regional cluster or in the same region as their
	   cluster's zone for zonal clusters. If unspecified, the cluster-level
	   `node_locations` will be used.

	   > Note: `node_locations` will not revert to the cluster's default set of zones
	   upon being unset. You must manually reconcile the list of zones with your
	   cluster.
	*/
	NodeLocations []string `json:"nodeLocations,omitempty" yaml:"nodeLocations,omitempty"`

	/*
	   Specify node upgrade settings to change how GKE upgrades nodes.
	   The maximum number of nodes upgraded simultaneously is limited to 20. Structure is documented below.
	*/
	UpgradeSettings types.Container_NodePoolUpgradeSettings `json:"upgradeSettings,omitempty" yaml:"upgradeSettings,omitempty"`

	/*
	   Configuration required by cluster autoscaler to adjust
	   the size of the node pool to the current cluster usage. Structure is documented below.
	*/
	Autoscaling types.Container_NodePoolAutoscaling `json:"autoscaling,omitempty" yaml:"autoscaling,omitempty"`

	/*
	   The name of the node pool. If left blank, the provider will
	   auto-generate a unique name.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Parameters used in creating the node pool. See
	   gcp.container.Cluster for schema.
	*/
	NodeConfig types.Container_NodePoolNodeConfig `json:"nodeConfig,omitempty" yaml:"nodeConfig,omitempty"`

	/*
	   Specifies a custom placement policy for the
	   nodes.
	*/
	PlacementPolicy types.Container_NodePoolPlacementPolicy `json:"placementPolicy,omitempty" yaml:"placementPolicy,omitempty"`

	/*
	   The Kubernetes version for the nodes in this pool. Note that if this field
	   and `auto_upgrade` are both specified, they will fight each other for what the node version should
	   be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
	   recommended that you specify explicit versions as the provider will see spurious diffs
	   when fuzzy versions are used. See the `gcp.container.getEngineVersions` data source's
	   `version_prefix` field to approximate fuzzy versions in a provider-compatible way.
	*/
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
