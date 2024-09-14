package container

import types "libds/gcp/types"

type Cluster struct {
	/*
	   Monitoring configuration for the cluster.
	   Structure is documented below.
	*/
	MonitoringConfig types.Container_ClusterMonitoringConfig `json:"monitoringConfig,omitempty" yaml:"monitoringConfig,omitempty"`

	/*
	   The monitoring service that the cluster
	   should write metrics to.
	   Automatically send metrics from pods in the cluster to the Google Cloud Monitoring API.
	   VM metrics will be collected by Google Compute Engine regardless of this setting
	   Available options include
	   `monitoring.googleapis.com`(Legacy Stackdriver), `monitoring.googleapis.com/kubernetes`(Stackdriver Kubernetes Engine Monitoring), and `none`.
	   Defaults to `monitoring.googleapis.com/kubernetes`
	*/
	MonitoringService string `json:"monitoringService,omitempty" yaml:"monitoringService,omitempty"`

	/*
	   List of node pools associated with this cluster.
	   See gcp.container.NodePool for schema.
	   --Warning:-- node pools defined inside a cluster can't be changed (or added/removed) after
	   cluster creation without deleting and recreating the entire cluster. Unless you absolutely need the ability
	   to say "these are the _only_ node pools associated with this cluster", use the
	   gcp.container.NodePool resource instead of this property.
	*/
	NodePools []types.Container_ClusterNodePool `json:"nodePools,omitempty" yaml:"nodePools,omitempty"`

	// Whether L4ILB Subsetting is enabled for this cluster.
	EnableL4IlbSubsetting bool `json:"enableL4IlbSubsetting,omitempty" yaml:"enableL4IlbSubsetting,omitempty"`

	// Configuration for [GKE Gateway API controller](https://cloud.google.com/kubernetes-engine/docs/concepts/gateway-api). Structure is documented below.
	GatewayApiConfig types.Container_ClusterGatewayApiConfig `json:"gatewayApiConfig,omitempty" yaml:"gatewayApiConfig,omitempty"`

	// Whether multi-networking is enabled for this cluster.
	EnableMultiNetworking bool `json:"enableMultiNetworking,omitempty" yaml:"enableMultiNetworking,omitempty"`

	/*
	   Configuration for the
	   [PodSecurityPolicy](https://cloud.google.com/kubernetes-engine/docs/how-to/pod-security-policies) feature.
	   Structure is documented below.
	*/
	PodSecurityPolicyConfig types.Container_ClusterPodSecurityPolicyConfig `json:"podSecurityPolicyConfig,omitempty" yaml:"podSecurityPolicyConfig,omitempty"`

	/*
	   The default maximum number of pods
	   per node in this cluster. This doesn't work on "routes-based" clusters, clusters
	   that don't have IP Aliasing enabled. See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
	   for more information.
	*/
	DefaultMaxPodsPerNode int `json:"defaultMaxPodsPerNode,omitempty" yaml:"defaultMaxPodsPerNode,omitempty"`

	/*
	   Whether to enable Kubernetes Alpha features for
	   this cluster. Note that when this option is enabled, the cluster cannot be upgraded
	   and will be automatically deleted after 30 days.
	*/
	EnableKubernetesAlpha bool `json:"enableKubernetesAlpha,omitempty" yaml:"enableKubernetesAlpha,omitempty"`

	// Fleet configuration for the cluster. Structure is documented below.
	Fleet types.Container_ClusterFleet `json:"fleet,omitempty" yaml:"fleet,omitempty"`

	/*
	   The authentication information for accessing the
	   Kubernetes master. Some values in this block are only returned by the API if
	   your service account has permission to get credentials for your GKE cluster. If
	   you see an unexpected diff unsetting your client cert, ensure you have the
	   `container.clusters.getCredentials` permission.
	   Structure is documented below.
	*/
	MasterAuth types.Container_ClusterMasterAuth `json:"masterAuth,omitempty" yaml:"masterAuth,omitempty"`

	/*
	   The name of the cluster, unique within the project and
	   location.

	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Determines whether alias IPs or routes will be used for pod IPs in the cluster.
	   Options are `VPC_NATIVE` or `ROUTES`. `VPC_NATIVE` enables [IP aliasing](https://cloud.google.com/kubernetes-engine/docs/how-to/ip-aliases). Newly created clusters will default to `VPC_NATIVE`.
	*/
	NetworkingMode string `json:"networkingMode,omitempty" yaml:"networkingMode,omitempty"`

	// Enable/Disable Protect API features for the cluster. Structure is documented below.
	ProtectConfig types.Container_ClusterProtectConfig `json:"protectConfig,omitempty" yaml:"protectConfig,omitempty"`

	// Structure is documented below.
	ServiceExternalIpsConfig types.Container_ClusterServiceExternalIpsConfig `json:"serviceExternalIpsConfig,omitempty" yaml:"serviceExternalIpsConfig,omitempty"`

	// The desired datapath provider for this cluster. This is set to `LEGACY_DATAPATH` by default, which uses the IPTables-based kube-proxy implementation. Set to `ADVANCED_DATAPATH` to enable Dataplane v2.
	DatapathProvider string `json:"datapathProvider,omitempty" yaml:"datapathProvider,omitempty"`

	// Description of the cluster.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Vertical Pod Autoscaling automatically adjusts the resources of pods controlled by it.
	   Structure is documented below.
	*/
	VerticalPodAutoscaling types.Container_ClusterVerticalPodAutoscaling `json:"verticalPodAutoscaling,omitempty" yaml:"verticalPodAutoscaling,omitempty"`

	/*
	   The ID of the project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Enable/Disable Security Posture API features for the cluster. Structure is documented below.
	SecurityPostureConfig types.Container_ClusterSecurityPostureConfig `json:"securityPostureConfig,omitempty" yaml:"securityPostureConfig,omitempty"`

	/*
	   Configuration for [direct-path (via ALTS) with workload identity.](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1beta1/projects.locations.clusters#workloadaltsconfig). Structure is documented below.

	   <a name="nested_default_snat_status"></a>The `default_snat_status` block supports
	*/
	WorkloadAltsConfig types.Container_ClusterWorkloadAltsConfig `json:"workloadAltsConfig,omitempty" yaml:"workloadAltsConfig,omitempty"`

	/*
	   Configuration for the
	   [Cost Allocation](https://cloud.google.com/kubernetes-engine/docs/how-to/cost-allocations) feature.
	   Structure is documented below.
	*/
	CostManagementConfig types.Container_ClusterCostManagementConfig `json:"costManagementConfig,omitempty" yaml:"costManagementConfig,omitempty"`

	/*
	   Parameters used in creating the default node pool.
	   Generally, this field should not be used at the same time as a
	   `gcp.container.NodePool` or a `node_pool` block; this configuration
	   manages the default node pool, which isn't recommended to be used.
	   Structure is documented below.
	*/
	NodeConfig types.Container_ClusterNodeConfig `json:"nodeConfig,omitempty" yaml:"nodeConfig,omitempty"`

	// . Structure is documented below.
	IdentityServiceConfig types.Container_ClusterIdentityServiceConfig `json:"identityServiceConfig,omitempty" yaml:"identityServiceConfig,omitempty"`

	/*
	   The number of nodes to create in this
	   cluster's default node pool. In regional or multi-zonal clusters, this is the
	   number of nodes per zone. Must be set if `node_pool` is not set. If you're using
	   `gcp.container.NodePool` objects with no default node pool, you'll need to
	   set this to a value of at least `1`, alongside setting
	   `remove_default_node_pool` to `true`.
	*/
	InitialNodeCount int `json:"initialNodeCount,omitempty" yaml:"initialNodeCount,omitempty"`

	/*
	   The Kubernetes version on the nodes. Must either be unset
	   or set to the same value as `min_master_version` on create. Defaults to the default
	   version set by GKE which is not necessarily the latest version. This only affects
	   nodes in the default node pool. While a fuzzy version can be specified, it's
	   recommended that you specify explicit versions as the provider will see spurious diffs
	   when fuzzy versions are used. See the `gcp.container.getEngineVersions` data source's
	   `version_prefix` field to approximate fuzzy versions.
	   To update nodes in other node pools, use the `version` attribute on the node pool.
	*/
	NodeVersion string `json:"nodeVersion,omitempty" yaml:"nodeVersion,omitempty"`

	// The GCE resource labels (a map of key/value pairs) to be applied to the cluster.
	ResourceLabels map[string]string `json:"resourceLabels,omitempty" yaml:"resourceLabels,omitempty"`

	/*
	   Workload Identity allows Kubernetes service accounts to act as a user-managed
	   [Google IAM Service Account](https://cloud.google.com/iam/docs/service-accounts#user-managed_service_accounts).
	   Structure is documented below.
	*/
	WorkloadIdentityConfig types.Container_ClusterWorkloadIdentityConfig `json:"workloadIdentityConfig,omitempty" yaml:"workloadIdentityConfig,omitempty"`

	// Configuration for [Using Cloud DNS for GKE](https://cloud.google.com/kubernetes-engine/docs/how-to/cloud-dns). Structure is documented below.
	DnsConfig types.Container_ClusterDnsConfig `json:"dnsConfig,omitempty" yaml:"dnsConfig,omitempty"`

	/*
	   Configuration for Kubernetes Beta APIs.
	   Structure is documented below.
	*/
	EnableK8sBetaApis types.Container_ClusterEnableK8sBetaApis `json:"enableK8sBetaApis,omitempty" yaml:"enableK8sBetaApis,omitempty"`

	/*
	   Configuration options for the
	   [NetworkPolicy](https://kubernetes.io/docs/concepts/services-networking/networkpolicies/)
	   feature. Structure is documented below.
	*/
	NetworkPolicy types.Container_ClusterNetworkPolicy `json:"networkPolicy,omitempty" yaml:"networkPolicy,omitempty"`

	// Configuration for the [cluster upgrade notifications](https://cloud.google.com/kubernetes-engine/docs/how-to/cluster-upgrade-notifications) feature. Structure is documented below.
	NotificationConfig types.Container_ClusterNotificationConfig `json:"notificationConfig,omitempty" yaml:"notificationConfig,omitempty"`

	/*
	   Configuration for the
	   [ResourceUsageExportConfig](https://cloud.google.com/kubernetes-engine/docs/how-to/cluster-usage-metering) feature.
	   Structure is documented below.
	*/
	ResourceUsageExportConfig types.Container_ClusterResourceUsageExportConfig `json:"resourceUsageExportConfig,omitempty" yaml:"resourceUsageExportConfig,omitempty"`

	/*
	   Configuration options for the Binary
	   Authorization feature. Structure is documented below.
	*/
	BinaryAuthorization types.Container_ClusterBinaryAuthorization `json:"binaryAuthorization,omitempty" yaml:"binaryAuthorization,omitempty"`

	// Structure is documented below.
	DatabaseEncryption types.Container_ClusterDatabaseEncryption `json:"databaseEncryption,omitempty" yaml:"databaseEncryption,omitempty"`

	/*
	   Whether the ABAC authorizer is enabled for this cluster.
	   When enabled, identities in the system, including service accounts, nodes, and controllers,
	   will have statically granted permissions beyond those provided by the RBAC configuration or IAM.
	   Defaults to `false`
	*/
	EnableLegacyAbac bool `json:"enableLegacyAbac,omitempty" yaml:"enableLegacyAbac,omitempty"`

	/*
	   Logging configuration for the cluster.
	   Structure is documented below.
	*/
	LoggingConfig types.Container_ClusterLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`

	/*
	   The minimum version of the master. GKE
	   will auto-update the master to new versions, so this does not guarantee the
	   current master version--use the read-only `master_version` field to obtain that.
	   If unset, the cluster's version will be set by GKE to the version of the most recent
	   official release (which is not necessarily the latest version).  Most users will find
	   the `gcp.container.getEngineVersions` data source useful - it indicates which versions
	   are available. If you intend to specify versions manually,
	   [the docs](https://cloud.google.com/kubernetes-engine/versioning-and-upgrades#specifying_cluster_version)
	   describe the various acceptable formats for this field.

	   > If you are using the `gcp.container.getEngineVersions` datasource with a regional cluster, ensure that you have provided a `location`
	   to the datasource. A region can have a different set of supported versions than its corresponding zones, and not all zones in a
	   region are guaranteed to support the same version.
	*/
	MinMasterVersion string `json:"minMasterVersion,omitempty" yaml:"minMasterVersion,omitempty"`

	/*
	   The name or self_link of the Google Compute Engine
	   subnetwork in which the cluster's instances are launched.
	*/
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`

	/*
	   Whether or not to allow Terraform to destroy the instance. Defaults to true. Unless this field is set to false in
	   Terraform state, a terraform destroy or terraform apply that would delete the cluster will fail.
	*/
	DeletionProtection bool `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`

	// Whether Intra-node visibility is enabled for this cluster. This makes same node pod to pod traffic visible for VPC network.
	EnableIntranodeVisibility bool `json:"enableIntranodeVisibility,omitempty" yaml:"enableIntranodeVisibility,omitempty"`

	/*
	   Configuration of cluster IP allocation for
	   VPC-native clusters. If this block is unset during creation, it will be set by the GKE backend.
	   Structure is documented below.
	*/
	IpAllocationPolicy types.Container_ClusterIpAllocationPolicy `json:"ipAllocationPolicy,omitempty" yaml:"ipAllocationPolicy,omitempty"`

	/*
	   The maintenance policy to use for the cluster. Structure is
	   documented below.
	*/
	MaintenancePolicy types.Container_ClusterMaintenancePolicy `json:"maintenancePolicy,omitempty" yaml:"maintenancePolicy,omitempty"`

	/*
	   Enable Autopilot for this cluster. Defaults to `false`.
	   Note that when this option is enabled, certain features of Standard GKE are not available.
	   See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/concepts/autopilot-overview#comparison)
	   for available features.
	*/
	EnableAutopilot bool `json:"enableAutopilot,omitempty" yaml:"enableAutopilot,omitempty"`

	/*
	   The name or self_link of the Google Compute Engine
	   network to which the cluster is connected. For Shared VPC, set this to the self link of the
	   shared network.
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The configuration for addons supported by GKE.
	   Structure is documented below.
	*/
	AddonsConfig types.Container_ClusterAddonsConfig `json:"addonsConfig,omitempty" yaml:"addonsConfig,omitempty"`

	/*
	   Per-cluster configuration of Node Auto-Provisioning with Cluster Autoscaler to
	   automatically adjust the size of the cluster and create/delete node pools based
	   on the current needs of the cluster's workload. See the
	   [guide to using Node Auto-Provisioning](https://cloud.google.com/kubernetes-engine/docs/how-to/node-auto-provisioning)
	   for more details. Structure is documented below.
	*/
	ClusterAutoscaling types.Container_ClusterClusterAutoscaling `json:"clusterAutoscaling,omitempty" yaml:"clusterAutoscaling,omitempty"`

	// Structure is documented below.
	MeshCertificates types.Container_ClusterMeshCertificates `json:"meshCertificates,omitempty" yaml:"meshCertificates,omitempty"`

	/*
	   If `true`, deletes the default node
	   pool upon cluster creation. If you're using `gcp.container.NodePool`
	   resources with no default node pool, this should be set to `true`, alongside
	   setting `initial_node_count` to at least `1`.
	*/
	RemoveDefaultNodePool bool `json:"removeDefaultNodePool,omitempty" yaml:"removeDefaultNodePool,omitempty"`

	/*
	   The list of zones in which the cluster's nodes
	   are located. Nodes must be in the region of their regional cluster or in the
	   same region as their cluster's zone for zonal clusters. If this is specified for
	   a zonal cluster, omit the cluster's zone.

	   > A "multi-zonal" cluster is a zonal cluster with at least one additional zone
	   defined; in a multi-zonal cluster, the cluster master is only present in a
	   single zone while nodes are present in each of the primary zone and the node
	   locations. In contrast, in a regional cluster, cluster master nodes are present
	   in multiple zones in the region. For that reason, regional clusters should be
	   preferred.
	*/
	NodeLocations []string `json:"nodeLocations,omitempty" yaml:"nodeLocations,omitempty"`

	/*
	   Node pool configs that apply to auto-provisioned node pools in
	   [autopilot](https://cloud.google.com/kubernetes-engine/docs/concepts/autopilot-overview#comparison) clusters and
	   [node auto-provisioning](https://cloud.google.com/kubernetes-engine/docs/how-to/node-auto-provisioning)-enabled clusters. Structure is documented below.
	*/
	NodePoolAutoConfig types.Container_ClusterNodePoolAutoConfig `json:"nodePoolAutoConfig,omitempty" yaml:"nodePoolAutoConfig,omitempty"`

	/*
	   The location (region or zone) in which the cluster
	   master will be created, as well as the default node location. If you specify a
	   zone (such as `us-central1-a`), the cluster will be a zonal cluster with a
	   single cluster master. If you specify a region (such as `us-west1`), the
	   cluster will be a regional cluster with multiple masters spread across zones in
	   the region, and with default node locations in those zones as well
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Default NodePool settings for the entire cluster. These settings are overridden if specified on the specific NodePool object. Structure is documented below.
	NodePoolDefaults types.Container_ClusterNodePoolDefaults `json:"nodePoolDefaults,omitempty" yaml:"nodePoolDefaults,omitempty"`

	/*
	   Configuration for the
	   [Google Groups for GKE](https://cloud.google.com/kubernetes-engine/docs/how-to/role-based-access-control#groups-setup-gsuite) feature.
	   Structure is documented below.
	*/
	AuthenticatorGroupsConfig types.Container_ClusterAuthenticatorGroupsConfig `json:"authenticatorGroupsConfig,omitempty" yaml:"authenticatorGroupsConfig,omitempty"`

	// [GKE SNAT](https://cloud.google.com/kubernetes-engine/docs/how-to/ip-masquerade-agent#how_ipmasq_works) DefaultSnatStatus contains the desired state of whether default sNAT should be disabled on the cluster, [API doc](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1beta1/projects.locations.clusters#networkconfig). Structure is documented below
	DefaultSnatStatus types.Container_ClusterDefaultSnatStatus `json:"defaultSnatStatus,omitempty" yaml:"defaultSnatStatus,omitempty"`

	/*
	   Whether to enable Cloud TPU resources in this cluster.
	   See the [official documentation](https://cloud.google.com/tpu/docs/kubernetes-engine-setup).
	*/
	EnableTpu bool `json:"enableTpu,omitempty" yaml:"enableTpu,omitempty"`

	/*
	   The desired
	   configuration options for master authorized networks. Omit the
	   nested `cidr_blocks` attribute to disallow external access (except
	   the cluster node IPs, which GKE automatically whitelists).
	   Structure is documented below.
	*/
	MasterAuthorizedNetworksConfig types.Container_ClusterMasterAuthorizedNetworksConfig `json:"masterAuthorizedNetworksConfig,omitempty" yaml:"masterAuthorizedNetworksConfig,omitempty"`

	/*
	   Configuration options for the [Release channel](https://cloud.google.com/kubernetes-engine/docs/concepts/release-channels)
	   feature, which provide more control over automatic upgrades of your GKE clusters.
	   When updating this field, GKE imposes specific version requirements. See
	   [Selecting a new release channel](https://cloud.google.com/kubernetes-engine/docs/concepts/release-channels#selecting_a_new_release_channel)
	   for more details; the `gcp.container.getEngineVersions` datasource can provide
	   the default version for a channel. Note that removing the `release_channel`
	   field from your config will cause the provider to stop managing your cluster's
	   release channel, but will not unenroll it. Instead, use the `"UNSPECIFIED"`
	   channel. Structure is documented below.
	*/
	ReleaseChannel types.Container_ClusterReleaseChannel `json:"releaseChannel,omitempty" yaml:"releaseChannel,omitempty"`

	/*
	   The IP address range of the Kubernetes pods
	   in this cluster in CIDR notation (e.g. `10.96.0.0/14`). Leave blank to have one
	   automatically chosen or specify a `/14` block in `10.0.0.0/8`. This field will
	   default a new cluster to routes-based, where `ip_allocation_policy` is not defined.
	*/
	ClusterIpv4Cidr string `json:"clusterIpv4Cidr,omitempty" yaml:"clusterIpv4Cidr,omitempty"`

	// Whether FQDN Network Policy is enabled on this cluster. Users who enable this feature for existing Standard clusters must restart the GKE Dataplane V2 `anetd` DaemonSet after enabling it. See the [Enable FQDN Network Policy in an existing cluster](https://cloud.google.com/kubernetes-engine/docs/how-to/fqdn-network-policies#enable_fqdn_network_policy_in_an_existing_cluster) for more information.
	EnableFqdnNetworkPolicy bool `json:"enableFqdnNetworkPolicy,omitempty" yaml:"enableFqdnNetworkPolicy,omitempty"`

	/*
	   The logging service that the cluster should
	   write logs to. Available options include `logging.googleapis.com`(Legacy Stackdriver),
	   `logging.googleapis.com/kubernetes`(Stackdriver Kubernetes Engine Logging), and `none`. Defaults to `logging.googleapis.com/kubernetes`
	*/
	LoggingService string `json:"loggingService,omitempty" yaml:"loggingService,omitempty"`

	/*
	   Configuration for [private clusters](https://cloud.google.com/kubernetes-engine/docs/how-to/private-clusters),
	   clusters with private nodes. Structure is documented below.
	*/
	PrivateClusterConfig types.Container_ClusterPrivateClusterConfig `json:"privateClusterConfig,omitempty" yaml:"privateClusterConfig,omitempty"`

	// TPU configuration for the cluster.
	TpuConfig types.Container_ClusterTpuConfig `json:"tpuConfig,omitempty" yaml:"tpuConfig,omitempty"`

	/*
	   Configuration for
	   [ClusterTelemetry](https://cloud.google.com/monitoring/kubernetes-engine/installing#controlling_the_collection_of_application_logs) feature,
	   Structure is documented below.
	*/
	ClusterTelemetry types.Container_ClusterClusterTelemetry `json:"clusterTelemetry,omitempty" yaml:"clusterTelemetry,omitempty"`

	// Configuration for [Confidential Nodes](https://cloud.google.com/kubernetes-engine/docs/how-to/confidential-gke-nodes) feature. Structure is documented below documented below.
	ConfidentialNodes types.Container_ClusterConfidentialNodes `json:"confidentialNodes,omitempty" yaml:"confidentialNodes,omitempty"`

	// The desired state of IPv6 connectivity to Google Services. By default, no private IPv6 access to or from Google Services (all access will be via IPv4).
	PrivateIpv6GoogleAccess string `json:"privateIpv6GoogleAccess,omitempty" yaml:"privateIpv6GoogleAccess,omitempty"`

	/*
	   Enable NET_ADMIN for the cluster. Defaults to
	   `false`. This field should only be enabled for Autopilot clusters (`enable_autopilot`
	   set to `true`).
	*/
	AllowNetAdmin bool `json:"allowNetAdmin,omitempty" yaml:"allowNetAdmin,omitempty"`

	// Enable Shielded Nodes features on all nodes in this cluster.  Defaults to `true`.
	EnableShieldedNodes bool `json:"enableShieldedNodes,omitempty" yaml:"enableShieldedNodes,omitempty"`
}
