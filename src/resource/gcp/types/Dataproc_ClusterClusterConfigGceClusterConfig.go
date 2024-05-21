package types

type Dataproc_ClusterClusterConfigGceClusterConfig struct {
	/*
	   The GCP zone where your data is stored and used (i.e. where
	   the master and the worker nodes will be created in). If `region` is set to 'global' (default)
	   then `zone` is mandatory, otherwise GCP is able to make use of [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/auto-zone)
	   to determine this automatically for you.
	   Note: This setting additionally determines and restricts
	   which computing resources are available for use with other configs such as
	   `cluster_config.master_config.machine_type` and `cluster_config.worker_config.machine_type`.
	*/
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`

	/*
	   By default, clusters are not restricted to internal IP addresses,
	   and will have ephemeral external IP addresses assigned to each instance. If set to true, all
	   instances in the cluster will only have internal IP addresses. Note: Private Google Access
	   (also known as `privateIpGoogleAccess`) must be enabled on the subnetwork that the cluster
	   will be launched in.
	*/
	InternalIpOnly bool `json:"internalIpOnly,omitempty" yaml:"internalIpOnly,omitempty"`

	/*
	   The service account to be used by the Node VMs.
	   If not specified, the "default" service account is used.
	*/
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	/*
	   The set of Google API scopes
	   to be made available on all of the node VMs under the `service_account`
	   specified. Both OAuth2 URLs and gcloud
	   short names are supported. To allow full access to all Cloud APIs, use the
	   `cloud-platform` scope. See a complete list of scopes [here](https://cloud.google.com/sdk/gcloud/reference/alpha/compute/instances/set-scopes#--scopes).
	*/
	ServiceAccountScopes []string `json:"serviceAccountScopes,omitempty" yaml:"serviceAccountScopes,omitempty"`

	/*
	   The name or self_link of the Google Compute Engine
	   subnetwork the cluster will be part of. Conflicts with `network`.
	*/
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`

	/*
	   Shielded Instance Config for clusters using [Compute Engine Shielded VMs](https://cloud.google.com/security/shielded-cloud/shielded-vm).

	   - - -
	*/
	ShieldedInstanceConfig Dataproc_ClusterClusterConfigGceClusterConfigShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty" yaml:"shieldedInstanceConfig,omitempty"`

	/*
	   The list of instance tags applied to instances in the cluster.
	   Tags are used to identify valid sources or targets for network firewalls.
	*/
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   A map of the Compute Engine metadata entries to add to all instances
	   (see [Project and instance metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)).
	*/
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	/*
	   The name or self_link of the Google Compute Engine
	   network to the cluster will be part of. Conflicts with `subnetwork`.
	   If neither is specified, this defaults to the "default" network.
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	// Node Group Affinity for sole-tenant clusters.
	NodeGroupAffinity Dataproc_ClusterClusterConfigGceClusterConfigNodeGroupAffinity `json:"nodeGroupAffinity,omitempty" yaml:"nodeGroupAffinity,omitempty"`

	// Reservation Affinity for consuming zonal reservation.
	ReservationAffinity Dataproc_ClusterClusterConfigGceClusterConfigReservationAffinity `json:"reservationAffinity,omitempty" yaml:"reservationAffinity,omitempty"`
}
