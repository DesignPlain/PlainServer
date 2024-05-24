package types

type Codedeploy_DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess struct {
	// The action to take on instances in the original environment after a successful blue/green deployment.
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// The number of minutes to wait after a successful blue/green deployment before terminating instances from the original environment.
	TerminationWaitTimeInMinutes int `json:"terminationWaitTimeInMinutes,omitempty" yaml:"terminationWaitTimeInMinutes,omitempty"`
}
