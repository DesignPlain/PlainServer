package opensearch

import types "DesignSphere_Server/src/resource/aws/types"

type ServerlessVpcEndpoint struct {
	// Name of the interface endpoint.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// One or more security groups that define the ports, protocols, and sources for inbound traffic that you are authorizing into your endpoint. Up to 5 security groups can be provided.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// One or more subnet IDs from which you'll access OpenSearch Serverless. Up to 6 subnets can be provided.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	//
	Timeouts types.Opensearch_ServerlessVpcEndpointTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	/*
	   ID of the VPC from which you'll access OpenSearch Serverless.

	   The following arguments are optional:
	*/
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
