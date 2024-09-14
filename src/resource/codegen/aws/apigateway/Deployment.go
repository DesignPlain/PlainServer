package apigateway

import types "libds/aws/types"

type Deployment struct {
	// Description of the deployment
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// REST API identifier.
	RestApi string `json:"restApi,omitempty" yaml:"restApi,omitempty"`

	// Description to set on the stage managed by the `stage_name` argument.
	StageDescription string `json:"stageDescription,omitempty" yaml:"stageDescription,omitempty"`

	// Name of the stage to create with this deployment. If the specified stage already exists, it will be updated to point to the new deployment. We recommend using the `aws.apigateway.Stage` resource instead to manage stages.
	StageName string `json:"stageName,omitempty" yaml:"stageName,omitempty"`

	// Map of arbitrary keys and values that, when changed, will trigger a redeployment.
	Triggers map[string]string `json:"triggers,omitempty" yaml:"triggers,omitempty"`

	// Map to set on the stage managed by the `stage_name` argument.
	Variables map[string]string `json:"variables,omitempty" yaml:"variables,omitempty"`

	// Input configuration for the canary deployment when the deployment is a canary release deployment. See `canary_settings below.
	CanarySettings types.Apigateway_DeploymentCanarySettings `json:"canarySettings,omitempty" yaml:"canarySettings,omitempty"`
}
