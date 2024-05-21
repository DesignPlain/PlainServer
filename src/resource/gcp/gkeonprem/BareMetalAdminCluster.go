package gkeonprem

import types "DesignSphere_Server/src/resource/gcp/types"

type BareMetalAdminCluster struct {
	// The bare metal admin cluster name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Network configuration.
	   Structure is documented below.
	*/
	NetworkConfig types.Gkeonprem_BareMetalAdminClusterNetworkConfig `json:"networkConfig,omitempty" yaml:"networkConfig,omitempty"`

	/*
	   Specifies the workload node configurations.
	   Structure is documented below.
	*/
	NodeConfig types.Gkeonprem_BareMetalAdminClusterNodeConfig `json:"nodeConfig,omitempty" yaml:"nodeConfig,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The location of the resource.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Specifies the workload node configurations.
	   Structure is documented below.
	*/
	MaintenanceConfig types.Gkeonprem_BareMetalAdminClusterMaintenanceConfig `json:"maintenanceConfig,omitempty" yaml:"maintenanceConfig,omitempty"`

	/*
	   Specifies the cluster storage configuration.
	   Structure is documented below.
	*/
	Storage types.Gkeonprem_BareMetalAdminClusterStorage `json:"storage,omitempty" yaml:"storage,omitempty"`

	/*
	   Annotations on the Bare Metal Admin Cluster.
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
	   Specifies the control plane configuration.
	   Structure is documented below.
	*/
	ControlPlane types.Gkeonprem_BareMetalAdminClusterControlPlane `json:"controlPlane,omitempty" yaml:"controlPlane,omitempty"`

	/*
	   Specifies the cluster proxy configuration.
	   Structure is documented below.
	*/
	Proxy types.Gkeonprem_BareMetalAdminClusterProxy `json:"proxy,omitempty" yaml:"proxy,omitempty"`

	/*
	   Specifies the Admin Cluster's observability infrastructure.
	   Structure is documented below.
	*/
	ClusterOperations types.Gkeonprem_BareMetalAdminClusterClusterOperations `json:"clusterOperations,omitempty" yaml:"clusterOperations,omitempty"`

	// A human readable description of this Bare Metal Admin Cluster.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Specifies the node access related settings for the bare metal user cluster.
	   Structure is documented below.
	*/
	NodeAccessConfig types.Gkeonprem_BareMetalAdminClusterNodeAccessConfig `json:"nodeAccessConfig,omitempty" yaml:"nodeAccessConfig,omitempty"`

	/*
	   Specifies the security related settings for the Bare Metal User Cluster.
	   Structure is documented below.
	*/
	SecurityConfig types.Gkeonprem_BareMetalAdminClusterSecurityConfig `json:"securityConfig,omitempty" yaml:"securityConfig,omitempty"`

	// A human readable description of this Bare Metal Admin Cluster.
	BareMetalVersion string `json:"bareMetalVersion,omitempty" yaml:"bareMetalVersion,omitempty"`

	/*
	   Specifies the load balancer configuration.
	   Structure is documented below.
	*/
	LoadBalancer types.Gkeonprem_BareMetalAdminClusterLoadBalancer `json:"loadBalancer,omitempty" yaml:"loadBalancer,omitempty"`
}
