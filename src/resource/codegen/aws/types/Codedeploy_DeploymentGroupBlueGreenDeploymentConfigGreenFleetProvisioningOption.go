package types

type Codedeploy_DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption struct {
	/*
	   The method used to add instances to a replacement environment.
	   - `DISCOVER_EXISTING`: Use instances that already exist or will be created manually.
	   - `COPY_AUTO_SCALING_GROUP`: Use settings from a specified --Auto Scaling-- group to define and create instances in a new Auto Scaling group. _Exactly one Auto Scaling group must be specified_ when selecting `COPY_AUTO_SCALING_GROUP`. Use `autoscaling_groups` to specify the Auto Scaling group.
	*/
	Action string `json:"action,omitempty" yaml:"action,omitempty"`
}
