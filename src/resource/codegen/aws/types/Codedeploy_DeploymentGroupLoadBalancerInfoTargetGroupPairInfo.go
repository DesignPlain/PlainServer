package types

type Codedeploy_DeploymentGroupLoadBalancerInfoTargetGroupPairInfo struct {
	// Configuration block for the production traffic route (documented below).
	ProdTrafficRoute Codedeploy_DeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute `json:"prodTrafficRoute,omitempty" yaml:"prodTrafficRoute,omitempty"`

	// Configuration blocks for a target group within a target group pair (documented below).
	TargetGroups []Codedeploy_DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup `json:"targetGroups,omitempty" yaml:"targetGroups,omitempty"`

	// Configuration block for the test traffic route (documented below).
	TestTrafficRoute Codedeploy_DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute `json:"testTrafficRoute,omitempty" yaml:"testTrafficRoute,omitempty"`
}
