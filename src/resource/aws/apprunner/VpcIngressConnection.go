package apprunner

import types "DesignSphere_Server/src/resource/aws/types"

type VpcIngressConnection struct {
	// Specifications for the customer’s Amazon VPC and the related AWS PrivateLink VPC endpoint that are used to create the VPC Ingress Connection resource. See Ingress VPC Configuration below for more details.
	IngressVpcConfiguration types.Apprunner_VpcIngressConnectionIngressVpcConfiguration `json:"ingressVpcConfiguration,omitempty" yaml:"ingressVpcConfiguration,omitempty"`

	// A name for the VPC Ingress Connection resource. It must be unique across all the active VPC Ingress Connections in your AWS account in the AWS Region.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The Amazon Resource Name (ARN) for this App Runner service that is used to create the VPC Ingress Connection resource.
	ServiceArn string `json:"serviceArn,omitempty" yaml:"serviceArn,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
