package alb

import types "DesignSphere_Server/src/resource/aws/types"

type Listener struct {
	// ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `aws.lb.ListenerCertificate` resource.
	CertificateArn string `json:"certificateArn,omitempty" yaml:"certificateArn,omitempty"`

	// Configuration block for default actions. Detailed below.
	DefaultActions []types.Alb_ListenerDefaultAction `json:"defaultActions,omitempty" yaml:"defaultActions,omitempty"`

	// Name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
	SslPolicy string `json:"sslPolicy,omitempty" yaml:"sslPolicy,omitempty"`

	/*
	   A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.

	   > --NOTE::-- Please note that listeners that are attached to Application Load Balancers must use either `HTTP` or `HTTPS` protocols while listeners that are attached to Network Load Balancers must use the `TCP` protocol.
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Name of the Application-Layer Protocol Negotiation (ALPN) policy. Can be set if `protocol` is `TLS`. Valid values are `HTTP1Only`, `HTTP2Only`, `HTTP2Optional`, `HTTP2Preferred`, and `None`.
	AlpnPolicy string `json:"alpnPolicy,omitempty" yaml:"alpnPolicy,omitempty"`

	/*
	   ARN of the load balancer.

	   The following arguments are optional:
	*/
	LoadBalancerArn string `json:"loadBalancerArn,omitempty" yaml:"loadBalancerArn,omitempty"`

	// The mutual authentication configuration information. Detailed below.
	MutualAuthentication types.Alb_ListenerMutualAuthentication `json:"mutualAuthentication,omitempty" yaml:"mutualAuthentication,omitempty"`

	// Port on which the load balancer is listening. Not valid for Gateway Load Balancers.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Protocol for connections from clients to the load balancer. For Application Load Balancers, valid values are `HTTP` and `HTTPS`, with a default of `HTTP`. For Network Load Balancers, valid values are `TCP`, `TLS`, `UDP`, and `TCP_UDP`. Not valid to use `UDP` or `TCP_UDP` if dual-stack mode is enabled. Not valid for Gateway Load Balancers.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
}
