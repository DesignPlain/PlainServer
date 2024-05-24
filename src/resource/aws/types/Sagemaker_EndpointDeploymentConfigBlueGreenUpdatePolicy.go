package types

type Sagemaker_EndpointDeploymentConfigBlueGreenUpdatePolicy struct {
	// Maximum execution timeout for the deployment. Note that the timeout value should be larger than the total waiting time specified in `termination_wait_in_seconds` and `wait_interval_in_seconds`. Valid values are between `600` and `14400`.
	MaximumExecutionTimeoutInSeconds int `json:"maximumExecutionTimeoutInSeconds,omitempty" yaml:"maximumExecutionTimeoutInSeconds,omitempty"`

	// Additional waiting time in seconds after the completion of an endpoint deployment before terminating the old endpoint fleet. Default is `0`. Valid values are between `0` and `3600`.
	TerminationWaitInSeconds int `json:"terminationWaitInSeconds,omitempty" yaml:"terminationWaitInSeconds,omitempty"`

	// Defines the traffic routing strategy to shift traffic from the old fleet to the new fleet during an endpoint deployment. See Traffic Routing Configuration.
	TrafficRoutingConfiguration Sagemaker_EndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfiguration `json:"trafficRoutingConfiguration,omitempty" yaml:"trafficRoutingConfiguration,omitempty"`
}
