package types

type Dns_ManagedZonePrivateVisibilityConfigGkeCluster struct {
	/*
	   The resource name of the cluster to bind this ManagedZone to.
	   This should be specified in the format like
	   `projects/-/locations/-/clusters/-`
	*/
	GkeClusterName string `json:"gkeClusterName,omitempty" yaml:"gkeClusterName,omitempty"`
}
