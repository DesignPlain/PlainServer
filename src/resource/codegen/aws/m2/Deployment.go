package m2

import types "libds/aws/types"

type Deployment struct {
	// Start the application once deployed.
	Start bool `json:"start,omitempty" yaml:"start,omitempty"`

	//
	Timeouts types.M2_DeploymentTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Application to deploy.
	ApplicationId string `json:"applicationId,omitempty" yaml:"applicationId,omitempty"`

	// Version to application to deploy
	ApplicationVersion int `json:"applicationVersion,omitempty" yaml:"applicationVersion,omitempty"`

	// Environment to deploy application to.
	EnvironmentId string `json:"environmentId,omitempty" yaml:"environmentId,omitempty"`

	//
	ForceStop bool `json:"forceStop,omitempty" yaml:"forceStop,omitempty"`
}
