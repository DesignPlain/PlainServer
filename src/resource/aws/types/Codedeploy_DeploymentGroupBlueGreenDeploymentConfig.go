package types

type Codedeploy_DeploymentGroupBlueGreenDeploymentConfig struct {
	// Information about the action to take when newly provisioned instances are ready to receive traffic in a blue/green deployment (documented below).
	DeploymentReadyOption Codedeploy_DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption `json:"deploymentReadyOption,omitempty" yaml:"deploymentReadyOption,omitempty"`

	// Information about how instances are provisioned for a replacement environment in a blue/green deployment (documented below).
	GreenFleetProvisioningOption Codedeploy_DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption `json:"greenFleetProvisioningOption,omitempty" yaml:"greenFleetProvisioningOption,omitempty"`

	/*
	   Information about whether to terminate instances in the original fleet during a blue/green deployment (documented below).

	   _Only one `blue_green_deployment_config` is allowed_.
	*/
	TerminateBlueInstancesOnDeploymentSuccess Codedeploy_DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess `json:"terminateBlueInstancesOnDeploymentSuccess,omitempty" yaml:"terminateBlueInstancesOnDeploymentSuccess,omitempty"`
}
