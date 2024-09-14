package types

type Codedeploy_DeploymentGroupDeploymentStyle struct {
	// Indicates whether to route deployment traffic behind a load balancer. Valid Values are `WITH_TRAFFIC_CONTROL` or `WITHOUT_TRAFFIC_CONTROL`. Default is `WITHOUT_TRAFFIC_CONTROL`.
	DeploymentOption string `json:"deploymentOption,omitempty" yaml:"deploymentOption,omitempty"`

	/*
	   Indicates whether to run an in-place deployment or a blue/green deployment. Valid Values are `IN_PLACE` or `BLUE_GREEN`. Default is `IN_PLACE`.

	   _Only one `deployment_style` is allowed_.
	*/
	DeploymentType string `json:"deploymentType,omitempty" yaml:"deploymentType,omitempty"`
}
