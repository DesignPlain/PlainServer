package types

type Codedeploy_DeploymentConfigTrafficRoutingConfigTimeBasedLinear struct {
	// The number of minutes between each incremental traffic shift of a `TimeBasedLinear` deployment.
	Interval int `json:"interval,omitempty" yaml:"interval,omitempty"`

	// The percentage of traffic that is shifted at the start of each increment of a `TimeBasedLinear` deployment.
	Percentage int `json:"percentage,omitempty" yaml:"percentage,omitempty"`
}
