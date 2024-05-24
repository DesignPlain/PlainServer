package types

type Codedeploy_DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption struct {
	// When to reroute traffic from an original environment to a replacement environment in a blue/green deployment.
	ActionOnTimeout string `json:"actionOnTimeout,omitempty" yaml:"actionOnTimeout,omitempty"`

	// The number of minutes to wait before the status of a blue/green deployment changed to Stopped if rerouting is not started manually. Applies only to the `STOP_DEPLOYMENT` option for `action_on_timeout`.
	WaitTimeInMinutes int `json:"waitTimeInMinutes,omitempty" yaml:"waitTimeInMinutes,omitempty"`
}
