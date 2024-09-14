package types

type Codedeploy_DeploymentConfigTrafficRoutingConfigTimeBasedCanary struct {
	// The number of minutes between the first and second traffic shifts of a `TimeBasedCanary` deployment.
	Interval int `json:"interval,omitempty" yaml:"interval,omitempty"`

	// The percentage of traffic to shift in the first increment of a `TimeBasedCanary` deployment.
	Percentage int `json:"percentage,omitempty" yaml:"percentage,omitempty"`
}
