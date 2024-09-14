package types

type Container_AwsClusterNetworking struct {
	// Disable the per node pool subnet security group rules on the control plane security group. When set to true, you must also provide one or more security groups that ensure node pools are able to send requests to the control plane on TCP/443 and TCP/8132. Failure to do so may result in unavailable node pools.
	PerNodePoolSgRulesDisabled bool `json:"perNodePoolSgRulesDisabled,omitempty" yaml:"perNodePoolSgRulesDisabled,omitempty"`

	// All pods in the cluster are assigned an RFC1918 IPv4 address from these ranges. Only a single range is supported. This field cannot be changed after creation.
	PodAddressCidrBlocks []string `json:"podAddressCidrBlocks,omitempty" yaml:"podAddressCidrBlocks,omitempty"`

	// All services in the cluster are assigned an RFC1918 IPv4 address from these ranges. Only a single range is supported. This field cannot be changed after creation.
	ServiceAddressCidrBlocks []string `json:"serviceAddressCidrBlocks,omitempty" yaml:"serviceAddressCidrBlocks,omitempty"`

	/*
	   The VPC associated with the cluster. All component clusters (i.e. control plane and node pools) run on a single VPC. This field cannot be changed after creation.

	   - - -
	*/
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
