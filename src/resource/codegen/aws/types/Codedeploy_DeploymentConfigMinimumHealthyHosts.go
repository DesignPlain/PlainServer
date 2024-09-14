package types

type Codedeploy_DeploymentConfigMinimumHealthyHosts struct {
	/*
	   The value when the type is `FLEET_PERCENT` represents the minimum number of healthy instances as
	   a percentage of the total number of instances in the deployment. If you specify FLEET_PERCENT, at the start of the
	   deployment, AWS CodeDeploy converts the percentage to the equivalent number of instance and rounds up fractional instances.
	   When the type is `HOST_COUNT`, the value represents the minimum number of healthy instances as an absolute value.
	*/
	Value int `json:"value,omitempty" yaml:"value,omitempty"`

	// The type can either be `FLEET_PERCENT` or `HOST_COUNT`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
