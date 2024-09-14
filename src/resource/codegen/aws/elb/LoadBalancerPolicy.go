package elb

import types "libds/aws/types"

type LoadBalancerPolicy struct {
	// The name of the load balancer policy.
	PolicyName string `json:"policyName,omitempty" yaml:"policyName,omitempty"`

	// The policy type.
	PolicyTypeName string `json:"policyTypeName,omitempty" yaml:"policyTypeName,omitempty"`

	// The load balancer on which the policy is defined.
	LoadBalancerName string `json:"loadBalancerName,omitempty" yaml:"loadBalancerName,omitempty"`

	// Policy attribute to apply to the policy.
	PolicyAttributes []types.Elb_LoadBalancerPolicyPolicyAttribute `json:"policyAttributes,omitempty" yaml:"policyAttributes,omitempty"`
}
