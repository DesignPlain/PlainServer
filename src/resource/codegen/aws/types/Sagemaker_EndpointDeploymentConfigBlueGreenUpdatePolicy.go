package types

type Sagemaker_EndpointDeploymentConfigBlueGreenUpdatePolicy struct {
	//
	MaximumExecutionTimeoutInSeconds int `json:"maximumExecutionTimeoutInSeconds,omitempty" yaml:"maximumExecutionTimeoutInSeconds,omitempty"`

	// Additional waiting time in seconds after the completion of an endpoint deployment before terminating the old endpoint fleet. Default is `0`. Valid values are between `0` and `3600`.
	TerminationWaitInSeconds int `json:"terminationWaitInSeconds,omitempty" yaml:"terminationWaitInSeconds,omitempty"`

	// Defines the traffic routing strategy to shift traffic from the old fleet to the new fleet during an endpoint deployment. See Traffic Routing Configuration.
	TrafficRoutingConfiguration Sagemaker_EndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfiguration `json:"trafficRoutingConfiguration,omitempty" yaml:"trafficRoutingConfiguration,omitempty"`
}
