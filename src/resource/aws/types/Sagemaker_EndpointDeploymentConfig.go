package types

type Sagemaker_EndpointDeploymentConfig struct {
	// Automatic rollback configuration for handling endpoint deployment failures and recovery. See Auto Rollback Configuration.
	AutoRollbackConfiguration Sagemaker_EndpointDeploymentConfigAutoRollbackConfiguration `json:"autoRollbackConfiguration,omitempty" yaml:"autoRollbackConfiguration,omitempty"`

	// Update policy for a blue/green deployment. If this update policy is specified, SageMaker creates a new fleet during the deployment while maintaining the old fleet. SageMaker flips traffic to the new fleet according to the specified traffic routing configuration. Only one update policy should be used in the deployment configuration. If no update policy is specified, SageMaker uses a blue/green deployment strategy with all at once traffic shifting by default. See Blue Green Update Config.
	BlueGreenUpdatePolicy Sagemaker_EndpointDeploymentConfigBlueGreenUpdatePolicy `json:"blueGreenUpdatePolicy,omitempty" yaml:"blueGreenUpdatePolicy,omitempty"`

	// Specifies a rolling deployment strategy for updating a SageMaker endpoint. See Rolling Update Policy.
	RollingUpdatePolicy Sagemaker_EndpointDeploymentConfigRollingUpdatePolicy `json:"rollingUpdatePolicy,omitempty" yaml:"rollingUpdatePolicy,omitempty"`
}
