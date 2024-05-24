package elasticsearch

import types "DesignSphere_Server/src/resource/aws/types"

type VpcEndpoint struct {
	// Options to specify the subnets and security groups for the endpoint.
	VpcOptions types.Elasticsearch_VpcEndpointVpcOptions `json:"vpcOptions,omitempty" yaml:"vpcOptions,omitempty"`

	// Specifies the Amazon Resource Name (ARN) of the domain to create the endpoint for
	DomainArn string `json:"domainArn,omitempty" yaml:"domainArn,omitempty"`
}
