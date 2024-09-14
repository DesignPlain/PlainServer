package types

type Codedeploy_DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption struct {
	// The number of minutes to wait before the status of a blue/green deployment changed to Stopped if rerouting is not started manually. Applies only to the `STOP_DEPLOYMENT` option for `action_on_timeout`.
	WaitTimeInMinutes int `json:"waitTimeInMinutes,omitempty" yaml:"waitTimeInMinutes,omitempty"`

	/*
	   When to reroute traffic from an original environment to a replacement environment in a blue/green deployment.
	   - `CONTINUE_DEPLOYMENT`: Register new instances with the load balancer immediately after the new application revision is installed on the instances in the replacement environment.
	   - `STOP_DEPLOYMENT`: Do not register new instances with load balancer unless traffic is rerouted manually. If traffic is not rerouted manually before the end of the specified wait period, the deployment status is changed to Stopped.
	*/
	ActionOnTimeout string `json:"actionOnTimeout,omitempty" yaml:"actionOnTimeout,omitempty"`
}
