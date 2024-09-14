package apigatewayv2

import types "libds/aws/types"

type Stage struct {
	// Whether updates to an API automatically trigger a new deployment. Defaults to `false`. Applicable for HTTP APIs.
	AutoDeploy bool `json:"autoDeploy,omitempty" yaml:"autoDeploy,omitempty"`

	// Default route settings for the stage.
	DefaultRouteSettings types.Apigatewayv2_StageDefaultRouteSettings `json:"defaultRouteSettings,omitempty" yaml:"defaultRouteSettings,omitempty"`

	// Description for the stage. Must be less than or equal to 1024 characters in length.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Map that defines the stage variables for the stage.
	StageVariables map[string]string `json:"stageVariables,omitempty" yaml:"stageVariables,omitempty"`

	/*
	   Settings for logging access in this stage.
	   Use the `aws.apigateway.Account` resource to configure [permissions for CloudWatch Logging](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html#set-up-access-logging-permissions).
	*/
	AccessLogSettings types.Apigatewayv2_StageAccessLogSettings `json:"accessLogSettings,omitempty" yaml:"accessLogSettings,omitempty"`

	/*
	   Identifier of a client certificate for the stage. Use the `aws.apigateway.ClientCertificate` resource to configure a client certificate.
	   Supported only for WebSocket APIs.
	*/
	ClientCertificateId string `json:"clientCertificateId,omitempty" yaml:"clientCertificateId,omitempty"`

	// Deployment identifier of the stage. Use the `aws.apigatewayv2.Deployment` resource to configure a deployment.
	DeploymentId string `json:"deploymentId,omitempty" yaml:"deploymentId,omitempty"`

	/*
	   Name of the stage. Must be between 1 and 128 characters in length.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Route settings for the stage.
	RouteSettings []types.Apigatewayv2_StageRouteSetting `json:"routeSettings,omitempty" yaml:"routeSettings,omitempty"`

	// Map of tags to assign to the stage. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// API identifier.
	ApiId string `json:"apiId,omitempty" yaml:"apiId,omitempty"`
}
