package gkeonprem

import types "libds/gcp/types"

type VMwareCluster struct {
	/*
	   VMware User Cluster control plane nodes must have either 1 or 3 replicas.
	   Structure is documented below.
	*/
	ControlPlaneNode types.Gkeonprem_VMwareClusterControlPlaneNode `json:"controlPlaneNode,omitempty" yaml:"controlPlaneNode,omitempty"`

	// Enable control plane V2. Default to false.
	EnableControlPlaneV2 bool `json:"enableControlPlaneV2,omitempty" yaml:"enableControlPlaneV2,omitempty"`

	/*
	   Load Balancer configuration.
	   Structure is documented below.
	*/
	LoadBalancer types.Gkeonprem_VMwareClusterLoadBalancer `json:"loadBalancer,omitempty" yaml:"loadBalancer,omitempty"`

	// The Anthos clusters on the VMware version for your user cluster.
	OnPremVersion string `json:"onPremVersion,omitempty" yaml:"onPremVersion,omitempty"`

	/*
	   Storage configuration.
	   Structure is documented below.
	*/
	Storage types.Gkeonprem_VMwareClusterStorage `json:"storage,omitempty" yaml:"storage,omitempty"`

	/*
	   AAGConfig specifies whether to spread VMware User Cluster nodes across at
	   least three physical hosts in the datacenter.
	   Structure is documented below.
	*/
	AntiAffinityGroups types.Gkeonprem_VMwareClusterAntiAffinityGroups `json:"antiAffinityGroups,omitempty" yaml:"antiAffinityGroups,omitempty"`

	// The location of the resource.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Specifies upgrade policy for the cluster.
	   Structure is documented below.
	*/
	UpgradePolicy types.Gkeonprem_VMwareClusterUpgradePolicy `json:"upgradePolicy,omitempty" yaml:"upgradePolicy,omitempty"`

	// Enable VM tracking.
	VmTrackingEnabled bool `json:"vmTrackingEnabled,omitempty" yaml:"vmTrackingEnabled,omitempty"`

	/*
	   RBAC policy that will be applied and managed by GKE On-Prem.
	   Structure is documented below.
	*/
	Authorization types.Gkeonprem_VMwareClusterAuthorization `json:"authorization,omitempty" yaml:"authorization,omitempty"`

	/*
	   Configuration for auto repairing.
	   Structure is documented below.
	*/
	AutoRepairConfig types.Gkeonprem_VMwareClusterAutoRepairConfig `json:"autoRepairConfig,omitempty" yaml:"autoRepairConfig,omitempty"`

	/*
	   VmwareDataplaneV2Config specifies configuration for Dataplane V2.
	   Structure is documented below.
	*/
	DataplaneV2 types.Gkeonprem_VMwareClusterDataplaneV2 `json:"dataplaneV2,omitempty" yaml:"dataplaneV2,omitempty"`

	// A human readable description of this VMware User Cluster.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The admin cluster this VMware User Cluster belongs to.
	   This is the full resource name of the admin cluster's hub membership.
	   In the future, references to other resource types might be allowed if
	   admin clusters are modeled as their own resources.
	*/
	AdminClusterMembership string `json:"adminClusterMembership,omitempty" yaml:"adminClusterMembership,omitempty"`

	/*
	   Annotations on the VMware User Cluster.
	   This field has the same restrictions as Kubernetes annotations.
	   The total size of all keys and values combined is limited to 256k.
	   Key can have 2 segments: prefix (optional) and name (required),
	   separated by a slash (/).
	   Prefix must be a DNS subdomain.
	   Name must be 63 characters or less, begin and end with alphanumerics,
	   with dashes (-), underscores (_), dots (.), and alphanumerics between.

	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	// The VMware cluster name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The VMware User Cluster network configuration.
	   Structure is documented below.
	*/
	NetworkConfig types.Gkeonprem_VMwareClusterNetworkConfig `json:"networkConfig,omitempty" yaml:"networkConfig,omitempty"`

	/*
	   VmwareVCenterConfig specifies vCenter config for the user cluster.
	   Inherited from the admin cluster.
	   Structure is documented below.
	*/
	Vcenters []types.Gkeonprem_VMwareClusterVcenter `json:"vcenters,omitempty" yaml:"vcenters,omitempty"`
}
