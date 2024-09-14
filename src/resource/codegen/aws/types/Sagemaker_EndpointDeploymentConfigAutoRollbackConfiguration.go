package types

type Sagemaker_EndpointDeploymentConfigAutoRollbackConfiguration struct {
	// List of CloudWatch alarms in your account that are configured to monitor metrics on an endpoint. If any alarms are tripped during a deployment, SageMaker rolls back the deployment. See Alarms.
	Alarms []Sagemaker_EndpointDeploymentConfigAutoRollbackConfigurationAlarm `json:"alarms,omitempty" yaml:"alarms,omitempty"`
}
