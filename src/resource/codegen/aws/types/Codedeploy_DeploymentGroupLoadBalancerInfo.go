package types

type Codedeploy_DeploymentGroupLoadBalancerInfo struct {
	// The (Application/Network Load Balancer) target group to use in a deployment. Conflicts with `elb_info` and `target_group_pair_info`.
	TargetGroupInfos []Codedeploy_DeploymentGroupLoadBalancerInfoTargetGroupInfo `json:"targetGroupInfos,omitempty" yaml:"targetGroupInfos,omitempty"`

	// The (Application/Network Load Balancer) target group pair to use in a deployment. Conflicts with `elb_info` and `target_group_info`.
	TargetGroupPairInfo Codedeploy_DeploymentGroupLoadBalancerInfoTargetGroupPairInfo `json:"targetGroupPairInfo,omitempty" yaml:"targetGroupPairInfo,omitempty"`

	// The Classic Elastic Load Balancer to use in a deployment. Conflicts with `target_group_info` and `target_group_pair_info`.
	ElbInfos []Codedeploy_DeploymentGroupLoadBalancerInfoElbInfo `json:"elbInfos,omitempty" yaml:"elbInfos,omitempty"`
}
