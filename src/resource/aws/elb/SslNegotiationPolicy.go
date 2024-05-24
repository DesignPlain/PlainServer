package elb

import types "DesignSphere_Server/src/resource/aws/types"

type SslNegotiationPolicy struct {
	/*
	   The load balancer port to which the policy
	   should be applied. This must be an active listener on the load
	   balancer.
	*/
	LbPort int `json:"lbPort,omitempty" yaml:"lbPort,omitempty"`

	/*
	   The load balancer to which the policy
	   should be attached.
	*/
	LoadBalancer string `json:"loadBalancer,omitempty" yaml:"loadBalancer,omitempty"`

	// The name of the attribute
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Map of arbitrary keys and values that, when changed, will trigger a redeployment.

	   To set your attributes, please see the [AWS Elastic Load Balancing Developer Guide](http://docs.aws.amazon.com/ElasticLoadBalancing/latest/DeveloperGuide/elb-security-policy-table.html) for a listing of the supported SSL protocols, SSL options, and SSL ciphers.

	   > --NOTE:-- The AWS documentation references Server Order Preference, which the AWS Elastic Load Balancing API refers to as `Server-Defined-Cipher-Order`. If you wish to set Server Order Preference, use this value instead.
	*/
	Triggers map[string]string `json:"triggers,omitempty" yaml:"triggers,omitempty"`

	// An SSL Negotiation policy attribute. Each has two properties:
	Attributes []types.Elb_SslNegotiationPolicyAttribute `json:"attributes,omitempty" yaml:"attributes,omitempty"`
}
