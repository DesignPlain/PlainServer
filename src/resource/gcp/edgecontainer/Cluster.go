package edgecontainer

import types "DesignSphere_Server/src/resource/gcp/types"

type Cluster struct {
	/*
	   The configuration of the cluster control plane.
	   Structure is documented below.
	*/
	ControlPlane types.Edgecontainer_ClusterControlPlane `json:"controlPlane,omitempty" yaml:"controlPlane,omitempty"`

	/*
	   Remote control plane disk encryption options. This field is only used when
	   enabling CMEK support.
	   Structure is documented below.
	*/
	ControlPlaneEncryption types.Edgecontainer_ClusterControlPlaneEncryption `json:"controlPlaneEncryption,omitempty" yaml:"controlPlaneEncryption,omitempty"`

	// The target cluster version. For example: "1.5.0".
	TargetVersion string `json:"targetVersion,omitempty" yaml:"targetVersion,omitempty"`

	/*
	   RBAC policy that will be applied and managed by GEC.
	   Structure is documented below.
	*/
	Authorization types.Edgecontainer_ClusterAuthorization `json:"authorization,omitempty" yaml:"authorization,omitempty"`

	// The GDCE cluster name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The release channel a cluster is subscribed to.
	   Possible values are: `RELEASE_CHANNEL_UNSPECIFIED`, `NONE`, `REGULAR`.
	*/
	ReleaseChannel string `json:"releaseChannel,omitempty" yaml:"releaseChannel,omitempty"`

	/*
	   Config that customers are allowed to define for GDCE system add-ons.
	   Structure is documented below.
	*/
	SystemAddonsConfig types.Edgecontainer_ClusterSystemAddonsConfig `json:"systemAddonsConfig,omitempty" yaml:"systemAddonsConfig,omitempty"`

	// The location of the resource.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Address pools for cluster data plane external load balancing.
	ExternalLoadBalancerIpv4AddressPools []string `json:"externalLoadBalancerIpv4AddressPools,omitempty" yaml:"externalLoadBalancerIpv4AddressPools,omitempty"`

	/*
	   Fleet related configuration.
	   Fleets are a Google Cloud concept for logically organizing clusters,
	   letting you use and manage multi-cluster capabilities and apply
	   consistent policies across your systems.
	   Structure is documented below.
	*/
	Networking types.Edgecontainer_ClusterNetworking `json:"networking,omitempty" yaml:"networking,omitempty"`

	/*
	   The default maximum number of pods per node used if a maximum value is not
	   specified explicitly for a node pool in this cluster. If unspecified, the
	   Kubernetes default value will be used.
	*/
	DefaultMaxPodsPerNode int `json:"defaultMaxPodsPerNode,omitempty" yaml:"defaultMaxPodsPerNode,omitempty"`

	/*
	   User-defined labels for the edgecloud cluster.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Cluster-wide maintenance policy configuration.
	   Structure is documented below.
	*/
	MaintenancePolicy types.Edgecontainer_ClusterMaintenancePolicy `json:"maintenancePolicy,omitempty" yaml:"maintenancePolicy,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Fleet related configuration.
	   Fleets are a Google Cloud concept for logically organizing clusters,
	   letting you use and manage multi-cluster capabilities and apply
	   consistent policies across your systems.
	   Structure is documented below.
	*/
	Fleet types.Edgecontainer_ClusterFleet `json:"fleet,omitempty" yaml:"fleet,omitempty"`
}
