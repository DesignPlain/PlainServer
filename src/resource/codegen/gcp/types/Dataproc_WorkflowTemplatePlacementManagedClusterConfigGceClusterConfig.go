package types

type Dataproc_WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig struct {
	// If true, all instances in the cluster will only have internal IP addresses. By default, clusters are not restricted to internal IP addresses, and will have ephemeral external IP addresses assigned to each instance. This `internal_ip_only` restriction can only be enabled for subnetwork enabled networks, and all off-cluster dependencies must be configured to be accessible without external IP addresses.
	InternalIpOnly bool `json:"internalIpOnly,omitempty" yaml:"internalIpOnly,omitempty"`

	// The Compute Engine metadata entries to add to all instances (see (https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)).
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// The Compute Engine network to be used for machine communications. Cannot be specified with subnetwork_uri. If neither `network_uri` nor `subnetwork_uri` is specified, the "default" network of the project is used, if it exists. Cannot be a "Custom Subnet Network" (see /regions/global/default` - `default`
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	// The (https://cloud.google.com/compute/docs/access/service-accounts#default_service_account) is used.
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	// The URIs of service account scopes to be included in Compute Engine instances. The following base set of scopes is always included: - https://www.googleapis.com/auth/cloud.useraccounts.readonly - https://www.googleapis.com/auth/devstorage.read_write - https://www.googleapis.com/auth/logging.write If no scopes are specified, the following defaults are also provided: - https://www.googleapis.com/auth/bigquery - https://www.googleapis.com/auth/bigtable.admin.table - https://www.googleapis.com/auth/bigtable.data - https://www.googleapis.com/auth/devstorage.full_control
	ServiceAccountScopes []string `json:"serviceAccountScopes,omitempty" yaml:"serviceAccountScopes,omitempty"`

	// Shielded Instance Config for clusters using [Compute Engine Shielded VMs](https://cloud.google.com/security/shielded-cloud/shielded-vm). Structure defined below.
	ShieldedInstanceConfig Dataproc_WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty" yaml:"shieldedInstanceConfig,omitempty"`

	// The Compute Engine tags to add to all instances (see (https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The zone where the Compute Engine cluster will be located. On a create request, it is required in the "global" region. If omitted in a non-global Dataproc region, the service will pick a zone in the corresponding Compute Engine region. On a get request, zone will always be present. A full URL, partial URI, or short name are valid. Examples: - `https://www.googleapis.com/compute/v1/projects/` - `us-central1-f`
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`

	// Node Group Affinity for sole-tenant clusters.
	NodeGroupAffinity Dataproc_WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity `json:"nodeGroupAffinity,omitempty" yaml:"nodeGroupAffinity,omitempty"`

	// The type of IPv6 access for a cluster. Possible values: PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED, INHERIT_FROM_SUBNETWORK, OUTBOUND, BIDIRECTIONAL
	PrivateIpv6GoogleAccess string `json:"privateIpv6GoogleAccess,omitempty" yaml:"privateIpv6GoogleAccess,omitempty"`

	// Reservation Affinity for consuming Zonal reservation.
	ReservationAffinity Dataproc_WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity `json:"reservationAffinity,omitempty" yaml:"reservationAffinity,omitempty"`

	// The Compute Engine subnetwork to be used for machine communications. Cannot be specified with network_uri. A full URL, partial URI, or short name are valid. Examples: - `https://www.googleapis.com/compute/v1/projects//regions/us-east1/subnetworks/sub0` - `sub0`
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`
}
