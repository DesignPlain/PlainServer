package types

type Container_getClusterAddonsConfig struct {
	// The status of the NodeLocal DNSCache addon. It is disabled by default. Set enabled = true to enable.
	DnsCacheConfigs []Container_getClusterAddonsConfigDnsCacheConfig `json:"dnsCacheConfigs,omitempty" yaml:"dnsCacheConfigs,omitempty"`

	// The status of the GCS Fuse CSI driver addon, which allows the usage of gcs bucket as volumes. Defaults to disabled; set enabled = true to enable.
	GcsFuseCsiDriverConfigs []Container_getClusterAddonsConfigGcsFuseCsiDriverConfig `json:"gcsFuseCsiDriverConfigs,omitempty" yaml:"gcsFuseCsiDriverConfigs,omitempty"`

	// The status of the Horizontal Pod Autoscaling addon, which increases or decreases the number of replica pods a replication controller has based on the resource usage of the existing pods. It ensures that a Heapster pod is running in the cluster, which is also used by the Cloud Monitoring service. It is enabled by default; set disabled = true to disable.
	HorizontalPodAutoscalings []Container_getClusterAddonsConfigHorizontalPodAutoscaling `json:"horizontalPodAutoscalings,omitempty" yaml:"horizontalPodAutoscalings,omitempty"`

	// The status of the HTTP (L7) load balancing controller addon, which makes it easy to set up HTTP load balancers for services in a cluster. It is enabled by default; set disabled = true to disable.
	HttpLoadBalancings []Container_getClusterAddonsConfigHttpLoadBalancing `json:"httpLoadBalancings,omitempty" yaml:"httpLoadBalancings,omitempty"`

	// Whether we should enable the network policy addon for the master. This must be enabled in order to enable network policy for the nodes. To enable this, you must also define a network_policy block, otherwise nothing will happen. It can only be disabled if the nodes already do not have network policies enabled. Defaults to disabled; set disabled = false to enable.
	NetworkPolicyConfigs []Container_getClusterAddonsConfigNetworkPolicyConfig `json:"networkPolicyConfigs,omitempty" yaml:"networkPolicyConfigs,omitempty"`

	// The status of the CloudRun addon. It is disabled by default. Set disabled = false to enable.
	CloudrunConfigs []Container_getClusterAddonsConfigCloudrunConfig `json:"cloudrunConfigs,omitempty" yaml:"cloudrunConfigs,omitempty"`

	// The of the Config Connector addon.
	ConfigConnectorConfigs []Container_getClusterAddonsConfigConfigConnectorConfig `json:"configConnectorConfigs,omitempty" yaml:"configConnectorConfigs,omitempty"`

	// Whether this cluster should enable the Google Compute Engine Persistent Disk Container Storage Interface (CSI) Driver. Set enabled = true to enable. The Compute Engine persistent disk CSI Driver is enabled by default on newly created clusters for the following versions: Linux clusters: GKE version 1.18.10-gke.2100 or later, or 1.19.3-gke.2100 or later.
	GcePersistentDiskCsiDriverConfigs []Container_getClusterAddonsConfigGcePersistentDiskCsiDriverConfig `json:"gcePersistentDiskCsiDriverConfigs,omitempty" yaml:"gcePersistentDiskCsiDriverConfigs,omitempty"`

	// The status of the Filestore CSI driver addon, which allows the usage of filestore instance as volumes. Defaults to disabled; set enabled = true to enable.
	GcpFilestoreCsiDriverConfigs []Container_getClusterAddonsConfigGcpFilestoreCsiDriverConfig `json:"gcpFilestoreCsiDriverConfigs,omitempty" yaml:"gcpFilestoreCsiDriverConfigs,omitempty"`

	// The status of the Backup for GKE Agent addon. It is disabled by default. Set enabled = true to enable.
	GkeBackupAgentConfigs []Container_getClusterAddonsConfigGkeBackupAgentConfig `json:"gkeBackupAgentConfigs,omitempty" yaml:"gkeBackupAgentConfigs,omitempty"`

	// The status of the Istio addon.
	IstioConfigs []Container_getClusterAddonsConfigIstioConfig `json:"istioConfigs,omitempty" yaml:"istioConfigs,omitempty"`

	// Configuration for the KALM addon, which manages the lifecycle of k8s. It is disabled by default; Set enabled = true to enable.
	KalmConfigs []Container_getClusterAddonsConfigKalmConfig `json:"kalmConfigs,omitempty" yaml:"kalmConfigs,omitempty"`
}
