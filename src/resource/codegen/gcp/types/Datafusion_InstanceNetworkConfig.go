package types

type Datafusion_InstanceNetworkConfig struct {
	/*
	   The IP range in CIDR notation to use for the managed Data Fusion instance
	   nodes. This range must not overlap with any other ranges used in the Data Fusion instance network.
	*/
	IpAllocation string `json:"ipAllocation,omitempty" yaml:"ipAllocation,omitempty"`

	/*
	   Name of the network in the project with which the tenant project
	   will be peered for executing pipelines. In case of shared VPC where the network resides in another host
	   project the network should specified in the form of projects/{host-project-id}/global/networks/{network}
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`
}
