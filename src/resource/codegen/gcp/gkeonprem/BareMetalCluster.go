package gkeonprem

import types "libds/gcp/types"

type BareMetalCluster struct {
	/*
	   Specifies the workload node configurations.
	   Structure is documented below.
	*/
	NodeConfig types.Gkeonprem_BareMetalClusterNodeConfig `json:"nodeConfig,omitempty" yaml:"nodeConfig,omitempty"`

	/*
	   Specifies the cluster storage configuration.
	   Structure is documented below.
	*/
	Storage types.Gkeonprem_BareMetalClusterStorage `json:"storage,omitempty" yaml:"storage,omitempty"`

	/*
	   Annotations on the Bare Metal User Cluster.
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

	/*
	   Specifies the User Cluster's observability infrastructure.
	   Structure is documented below.
	*/
	ClusterOperations types.Gkeonprem_BareMetalClusterClusterOperations `json:"clusterOperations,omitempty" yaml:"clusterOperations,omitempty"`

	/*
	   Specifies the workload node configurations.
	   Structure is documented below.
	*/
	MaintenanceConfig types.Gkeonprem_BareMetalClusterMaintenanceConfig `json:"maintenanceConfig,omitempty" yaml:"maintenanceConfig,omitempty"`

	// The bare metal cluster name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A human readable description of this Bare Metal User Cluster.
	BareMetalVersion string `json:"bareMetalVersion,omitempty" yaml:"bareMetalVersion,omitempty"`

	/*
	   Network configuration.
	   Structure is documented below.
	*/
	NetworkConfig types.Gkeonprem_BareMetalClusterNetworkConfig `json:"networkConfig,omitempty" yaml:"networkConfig,omitempty"`

	/*
	   Specifies the node access related settings for the bare metal user cluster.
	   Structure is documented below.
	*/
	NodeAccessConfig types.Gkeonprem_BareMetalClusterNodeAccessConfig `json:"nodeAccessConfig,omitempty" yaml:"nodeAccessConfig,omitempty"`

	/*
	   Specifies the cluster proxy configuration.
	   Structure is documented below.
	*/
	Proxy types.Gkeonprem_BareMetalClusterProxy `json:"proxy,omitempty" yaml:"proxy,omitempty"`

	/*
	   The Admin Cluster this Bare Metal User Cluster belongs to.
	   This is the full resource name of the Admin Cluster's hub membership.
	*/
	AdminClusterMembership string `json:"adminClusterMembership,omitempty" yaml:"adminClusterMembership,omitempty"`

	/*
	   Specifies the control plane configuration.
	   Structure is documented below.
	*/
	ControlPlane types.Gkeonprem_BareMetalClusterControlPlane `json:"controlPlane,omitempty" yaml:"controlPlane,omitempty"`

	/*
	   OS environment related configurations.
	   Structure is documented below.
	*/
	OsEnvironmentConfig types.Gkeonprem_BareMetalClusterOsEnvironmentConfig `json:"osEnvironmentConfig,omitempty" yaml:"osEnvironmentConfig,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Specifies the security related settings for the Bare Metal User Cluster.
	   Structure is documented below.
	*/
	SecurityConfig types.Gkeonprem_BareMetalClusterSecurityConfig `json:"securityConfig,omitempty" yaml:"securityConfig,omitempty"`

	/*
	   The cluster upgrade policy.
	   Structure is documented below.
	*/
	UpgradePolicy types.Gkeonprem_BareMetalClusterUpgradePolicy `json:"upgradePolicy,omitempty" yaml:"upgradePolicy,omitempty"`

	/*
	   Binary Authorization related configurations.
	   Structure is documented below.
	*/
	BinaryAuthorization types.Gkeonprem_BareMetalClusterBinaryAuthorization `json:"binaryAuthorization,omitempty" yaml:"binaryAuthorization,omitempty"`

	// A human readable description of this Bare Metal User Cluster.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Specifies the load balancer configuration.
	   Structure is documented below.
	*/
	LoadBalancer types.Gkeonprem_BareMetalClusterLoadBalancer `json:"loadBalancer,omitempty" yaml:"loadBalancer,omitempty"`

	// The location of the resource.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}
