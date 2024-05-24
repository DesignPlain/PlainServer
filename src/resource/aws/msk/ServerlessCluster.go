package msk

import types "DesignSphere_Server/src/resource/aws/types"

type ServerlessCluster struct {
	// Specifies client authentication information for the serverless cluster. See below.
	ClientAuthentication types.Msk_ServerlessClusterClientAuthentication `json:"clientAuthentication,omitempty" yaml:"clientAuthentication,omitempty"`

	// The name of the serverless cluster.
	ClusterName string `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// VPC configuration information. See below.
	VpcConfigs []types.Msk_ServerlessClusterVpcConfig `json:"vpcConfigs,omitempty" yaml:"vpcConfigs,omitempty"`
}
