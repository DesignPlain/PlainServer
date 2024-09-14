package types

type Alloydb_ClusterNetworkConfig struct {
	/*
	   The name of the allocated IP range for the private IP AlloyDB cluster. For example: "google-managed-services-default".
	   If set, the instance IPs for this cluster will be created in the allocated range.
	*/
	AllocatedIpRange string `json:"allocatedIpRange,omitempty" yaml:"allocatedIpRange,omitempty"`

	/*
	   The resource link for the VPC network in which cluster resources are created and from which they are accessible via Private IP. The network must belong to the same project as the cluster.
	   It is specified in the form: "projects/{projectNumber}/global/networks/{network_id}".
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`
}
