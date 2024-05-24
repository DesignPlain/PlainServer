package elb

import types "DesignSphere_Server/src/resource/aws/types"

type LoadBalancerPolicy struct {
	// The policy type.
	PolicyTypeName string `json:"policyTypeName,omitempty" yaml:"policyTypeName,omitempty"`

	// The load balancer on which the policy is defined.
	LoadBalancerName string `json:"loadBalancerName,omitempty" yaml:"loadBalancerName,omitempty"`

	// Policy attribute to apply to the policy.
	PolicyAttributes []types.Elb_LoadBalancerPolicyPolicyAttribute `json:"policyAttributes,omitempty" yaml:"policyAttributes,omitempty"`

	// The name of the load balancer policy.
	PolicyName string `json:"policyName,omitempty" yaml:"policyName,omitempty"`
}
