package apprunner

import types "DesignSphere_Server/src/resource/aws/types"

type Deployment struct {
	// The Amazon Resource Name (ARN) of the App Runner service to start the deployment for.
	ServiceArn string `json:"serviceArn,omitempty" yaml:"serviceArn,omitempty"`

	//
	Timeouts types.Apprunner_DeploymentTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
