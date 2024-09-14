package types

type Codedeploy_DeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute struct {
	// List of Amazon Resource Names (ARNs) of the load balancer listeners. Must contain exactly one listener ARN.
	ListenerArns []string `json:"listenerArns,omitempty" yaml:"listenerArns,omitempty"`
}
